package sys_gen

import (
	"bytes"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"text/template"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
	"xcore/common/xgorm"
	"xcore/common/xtemplate"
	"xcore/common/xtype/xbool"
	"xcore/common/xtype/xstring"
)

type ISysGenService interface {
	AddGenTable(c *gin.Context, params []string) error
	GetTableColumns(c *gin.Context, params string) ([]*model.SysGenTableColumn, error)
	GetColumService() xgorm.IServiceFunctions[model.SysGenTableColumn]
	CodePreview(c *gin.Context, param string) (GenTablePreview, error)
	GetDB() *gorm.DB
}

type SysGenService struct {
	tableService  xgorm.IServiceFunctions[model.SysGenTable]
	columnService xgorm.IServiceFunctions[model.SysGenTableColumn]
	query         *query.Query
	db            *gorm.DB
}

type TmplGenParam struct {
	Table   *model.SysGenTable
	Columns []*model.SysGenTableColumn
}

func (s SysGenService) CodePreview(c *gin.Context, param string) (GenTablePreview, error) {
	var result GenTablePreview
	var genParam TmplGenParam
	tableQuery := s.query.SysGenTable
	columnQuery := s.query.SysGenTableColumn
	if count, err := tableQuery.WithContext(c).Where(tableQuery.TableName_.Eq(param)).Count(); err != nil || count <= 0 {
		return result, xerror.NewErrCode(xerror.GEN_NOT_EXIST_ERROR)
	}
	tableInSql, err := tableQuery.WithContext(c).Where(tableQuery.TableName_.Eq(param)).First()
	if err != nil {
		return GenTablePreview{}, err
	}
	genParam.Table = tableInSql
	columnsInSql, err := columnQuery.WithContext(c).Where(columnQuery.TableName_.Eq(param)).Find()
	if err != nil {
		return GenTablePreview{}, err
	}
	genParam.Columns = columnsInSql
	tmplBasePath := `./public/resources/template`
	goTmpfs := []string{"/go/xxx.type.tmpl", "/go/xxx.router.tmpl", "/go/xxx.service.tmpl"}
	for i, tmpl := range goTmpfs {
		tmplEngine := xtemplate.GetTemplate(tmplBasePath+tmpl, template.FuncMap{})
		var buffers bytes.Buffer
		err = tmplEngine.Execute(&buffers, genParam)
		if err != nil {
			log.Fatalln(err)
			return GenTablePreview{}, err
		}
		switch i {
		case 0:
			result.GoType = string(buffers.Bytes())
			break
		case 1:
			result.GoRouter = string(buffers.Bytes())
			break
		case 2:
			result.GoService = string(buffers.Bytes())
			break
		}
	}
	return result, nil
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
