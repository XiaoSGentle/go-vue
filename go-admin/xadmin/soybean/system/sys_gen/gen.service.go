package sys_gen

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xgorm"
	"xcore/common/xtype/xbool"
	"xcore/common/xtype/xstring"
)

type ISysGenService interface {
	AddGenTable(c *gin.Context, params []string) error
	GetTableColumns(c *gin.Context, params string) ([]*model.SysGenTableColumn, error)
	GetColumService() xgorm.IServiceFunctions[model.SysGenTableColumn]
	GetDB() *gorm.DB
}

type SysGenService struct {
	tableService  xgorm.IServiceFunctions[model.SysGenTable]
	columnService xgorm.IServiceFunctions[model.SysGenTableColumn]
	query         *query.Query
	db            *gorm.DB
}

func (s SysGenService) GetTableColumns(c *gin.Context, params string) ([]*model.SysGenTableColumn, error) {
	columnQuery := query.Use(s.db).SysGenTableColumn
	find, err := columnQuery.WithContext(c).Where(columnQuery.TableName_.Eq(params)).Find()
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (s SysGenService) GetColumService() xgorm.IServiceFunctions[model.SysGenTableColumn] {
	return s.columnService
}

func (s SysGenService) AddGenTable(c *gin.Context, params []string) error {
	columnQuery := query.Use(s.db).SysGenTableColumn
	tableQuery := query.Use(s.db).SysGenTable
	operatorInfo := s.tableService.GetOperatorInfo(c)
	notAllowedEditColumns := []string{"id", "version", "create_uid", "create_by", "create_time", "update_time", "update_uid", "update_by", "delete_tag"}
	for _, table := range params {
		_, _ = columnQuery.WithContext(c).Where(columnQuery.TableName_.Eq(table)).Delete()
		_, _ = tableQuery.WithContext(c).Where(tableQuery.TableName_.Eq(table)).Delete()
		// 如果有表
		if s.db.Migrator().HasTable(table) {
			// 给个UUID
			var genTable = model.SysGenTable{
				TableName_:     table,
				TableComment:   "",
				UpperCamelCase: xstring.SnakeToUpperCamelCase(table),
				LowerCamelCase: xstring.SnakeToLowerCamelCase(table),
				RelativePath:   "/" + table,
				CheckToken:     "1",
				CheckAuth:      "1",
				AddLog:         "1",
				Remarks:        "",
				CreateUID:      operatorInfo.Uid,
				CreateBy:       operatorInfo.NickName,
				CreateTime:     time.Now(),
			}
			//operator_info_bind.AddUpInfo(&genTable, operator_info_bind.ADD, c)
			err := tableQuery.WithContext(c).Save(&genTable)
			if err != nil {
				return err
			}
			// 获取表字段信息
			types, err := s.db.Migrator().ColumnTypes(table)
			if err != nil {
				return err
			}
			// 遍历字段插入
			var insertColumns []*model.SysGenTableColumn
			for _, columnType := range types {
				// 名称
				columnName := columnType.Name()
				// 字段长度
				lengthNum, _ := columnType.Length()
				// 字段描述
				comment, _ := columnType.Comment()
				_column := model.SysGenTableColumn{

					TableName_:     table,
					UpperCamelCase: xstring.SnakeToUpperCamelCase(columnName),
					LowerCamelCase: xstring.SnakeToLowerCamelCase(columnName),
					SnakeCase:      columnName,
					Comment:        comment,
					Length:         strconv.FormatInt(lengthNum, 10),
					GoType:         xstring.GoTypeConversion(columnName),
					TsType:         xstring.TsTypeConversion(columnName),
					Required:       "2",
					BaseColumn:     xbool.BooleanTo(slice.Contain(notAllowedEditColumns, columnName), "1", "2"),
					IsQuery:        "2",
					QueryType:      "EQ",
					HTMLType:       "INPUT",
					DictType:       "",
					Sort:           0,
					Version:        0,
					CreateUID:      operatorInfo.Uid,
					CreateBy:       operatorInfo.NickName,
					CreateTime:     time.Now(),
					UpdateTime:     time.Time{},
				}
				insertColumns = append(insertColumns, &_column)
			}
			err = columnQuery.WithContext(c).Create(insertColumns...)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s SysGenService) GetDB() *gorm.DB {
	return s.db
}

func NewSysGenService(db *gorm.DB) ISysGenService {
	return &SysGenService{
		db:            db,
		query:         query.Use(db),
		tableService:  xgorm.InjectService[model.SysGenTable](db),
		columnService: xgorm.InjectService[model.SysGenTableColumn](db),
	}
}
