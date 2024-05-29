package sys_gen

import (
	"bytes"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
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
	CodePreview(c *gin.Context, param string) ([]GenTablePreview, error)
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

func (s SysGenService) CodePreview(c *gin.Context, param string) ([]GenTablePreview, error) {
	var result []GenTablePreview
	var genParam TmplGenParam
	tableQuery := s.query.SysGenTable
	columnQuery := s.query.SysGenTableColumn
	if count, err := tableQuery.WithContext(c).Where(tableQuery.TableName_.Eq(param)).Count(); err != nil || count <= 0 {
		return result, xerror.NewErrCode(xerror.GEN_NOT_EXIST_ERROR)
	}
	tableInSql, err := tableQuery.WithContext(c).Where(tableQuery.TableName_.Eq(param)).First()
	if err != nil {
		return result, err
	}
	genParam.Table = tableInSql
	columnsInSql, err := columnQuery.WithContext(c).Where(columnQuery.TableName_.Eq(param)).Find()
	if err != nil {
		return result, err
	}
	genParam.Columns = columnsInSql
	tmplBasePath := `./public/resources/template`
	goTmpfs := []string{"/go/xxx.type.go.tmpl", "/go/xxx.router.go.tmpl", "/go/xxx.service.go.tmpl"}
	for _, tmpl := range goTmpfs {
		tmplEngine := xtemplate.GetTemplate(tmplBasePath+tmpl, s.genFuncMap())
		var buffers bytes.Buffer
		err = tmplEngine.Execute(&buffers, genParam)
		if err != nil {
			log.Fatalln(err)
			return result, err
		}

		result = append(result, GenTablePreview{
			Lang:        "golang",
			FileName:    strings.ReplaceAll(strings.ReplaceAll(tmpl, "/go/xxx", tableInSql.ShortName), ".tmpl", ""),
			FileContent: buffers.String(),
		})
	}
	vueTmpfs := []string{"/vue/xxx.api.ts.tmpl", "/vue/xxx-table-action.vue.tmpl", "/vue/namespace.d.ts.tmpl", "/vue/index.vue.tmpl", "/vue/i18n.ts.tmpl"}

	for _, tmpl := range vueTmpfs {
		tmplEngine := xtemplate.GetTemplate(tmplBasePath+tmpl, s.genFuncMap())
		var buffers bytes.Buffer
		err = tmplEngine.Execute(&buffers, genParam)
		if err != nil {
			log.Fatalln(err)
			return result, err
		}
		result = append(result, GenTablePreview{
			Lang:        "typescript",
			FileName:    strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmpl, "xxx", tableInSql.ShortName), "/vue/", ""), ".tmpl", ""),
			FileContent: buffers.String(),
		})
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
				ShortName:      xstring.SnackLastName(table),
				UpperCamelCase: xstring.SnakeToUpperCamelCase(table),
				LowerCamelCase: xstring.SnakeToLowerCamelCase(table),
				RelativePath:   "/" + xstring.SnackToPath(table),
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

func (s SysGenService) genFuncMap() template.FuncMap {
	return template.FuncMap{
		"AppendQueryEnd": func(i *model.SysGenTableColumn, is []*model.SysGenTableColumn, sign string) string {
			filters := slice.Filter(is, func(index int, item *model.SysGenTableColumn) bool {
				return item.IsQuery == "1"
			})
			var index int
			slice.FindBy(filters, func(_index int, item *model.SysGenTableColumn) bool {
				if item.SnakeCase == i.SnakeCase {
					index = _index
					return true
				}
				return false
			})
			return xbool.BooleanTo(len(filters)-1 == index, "", "|")
		},
		"AppendAddEnd": func(i *model.SysGenTableColumn, is []*model.SysGenTableColumn, sign string) string {
			filters := slice.Filter(is, func(index int, item *model.SysGenTableColumn) bool {
				return item.IsAdd == "1"
			})
			var index int
			slice.FindBy(filters, func(_index int, item *model.SysGenTableColumn) bool {
				if item.SnakeCase == i.SnakeCase {
					index = _index
					return true
				}
				return false
			})
			return xbool.BooleanTo(len(filters)-1 == index, "", "|")
		},
		"SnakeToLowerCamelCase": func(s string) string {
			return xstring.SnakeToLowerCamelCase(s)
		},
		"GenI18nType": func(s string) string {
			return xstring.SnakeToLowerCamelCase(s)
		},
		"GetTsDefaultValue": func(s string) string {
			switch s {
			case "":
				return ""
			default:
				return ""
			}
		},
		"GenRuleKey": func(i *model.SysGenTableColumn, is []*model.SysGenTableColumn) string {
			filters := slice.Filter(is, func(index int, item *model.SysGenTableColumn) bool {
				return item.Required == "1"
			})
			var index int
			slice.FindBy(filters, func(_index int, item *model.SysGenTableColumn) bool {
				if item.SnakeCase == i.SnakeCase {
					index = _index
					return true
				}
				return false
			})
			return xbool.BooleanTo(len(filters)-1 == index, i.LowerCamelCase, i.LowerCamelCase+"|")
		},
		"GenEditHtmlItem": func(i *model.SysGenTableColumn, t *model.SysGenTable) string {

			var renderStars string
			switch i.HTMLType {
			case "INPUT":
				renderStars = fmt.Sprintf(`
		<NFormItem :label="$t('page.%s.%s')" %s>
          <NInput v-model:value="model.%s" :placeholder="$t('page.%s.from.%s')" />
        </NFormItem>`, t.ShortName, i.LowerCamelCase, xbool.BooleanTo(i.Required == "1", fmt.Sprintf(`path="%s"`, i.LowerCamelCase), ""), i.LowerCamelCase, t.ShortName, i.LowerCamelCase)

			case "INPUT_TEXTAREA":
				renderStars = fmt.Sprintf(`
		<NFormItem :label="$t('page.%s.%s')" %s>
          <NInput v-model:value="model.%s"  :placeholder="$t('page.%s.from.%s')" type="textarea" />
        </NFormItem>`, t.ShortName, i.LowerCamelCase, xbool.BooleanTo(i.Required == "1", fmt.Sprintf(`path="%s"`, i.LowerCamelCase), ""), i.LowerCamelCase, t.ShortName, i.LowerCamelCase)

			case "DICT":
				renderStars = fmt.Sprintf(`
		<NFormItem :label="$t('page.%s.%s')" %s>
          <SysDict v-model:value="model.%s" dict-key="%s"/>
        </NFormItem>`, t.ShortName, i.LowerCamelCase, xbool.BooleanTo(i.Required == "1",
					fmt.Sprintf(`path="%s"`, i.LowerCamelCase), ""),
					i.LowerCamelCase, i.DictCode)
			case "RADIO":
				renderStars = fmt.Sprintf(`
		<NFormItem :label="$t('page.%s.%s')" %s>
		  <NRadioGroup v-model:value="model.%s">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>`, t.ShortName, i.LowerCamelCase, xbool.BooleanTo(i.Required == "1",
					fmt.Sprintf(`path="%s"`, i.LowerCamelCase), ""),
					i.LowerCamelCase)
			}
			return renderStars
		},
		"GenColumn": func(i *model.SysGenTableColumn, t *model.SysGenTable) string {
			var renderStars string
			switch i.HTMLType {
			case "INPUT":
				renderStars = fmt.Sprintf(` row => <span>{row.%s}</span>`, i.LowerCamelCase)

			case "INPUT_TEXTAREA":
				renderStars = fmt.Sprintf(`row => row.%s && <SysDict selectValue={row.%s} type="show" dictKey="%s" />`, i.LowerCamelCase, i.LowerCamelCase, i.DictCode)

			case "DICT":
				renderStars = fmt.Sprintf(`row => row.%s && <SysDict selectValue={row.%s} type="show" dictKey="%s" />`, i.LowerCamelCase, i.LowerCamelCase, i.DictCode)
			}
			return fmt.Sprintf(`
			{
			  key: '%s',
			  title: $t('page.%s.%s'),
			  align: 'center',
			  render: %s
			},	
`, i.LowerCamelCase, t.ShortName, i.LowerCamelCase, renderStars)
		},
		"RenderNative": func(content string) string {
			return "{{ " + content + " }}"
		},
	}
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
