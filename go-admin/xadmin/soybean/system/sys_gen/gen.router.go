package sys_gen

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"xadmin/soybean/dao/model"
	"xcore/common/xgorm"
	"xcore/common/xresponse"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var SysGenTableGroup = xcore.Group("/system/table", NewSysGenHandler, regSysGenTableMangerRoute)
var SysGenColumGroup = xcore.Group("/system/column", NewSysGenHandler, regSysGenColumMangerRoute)

func regSysGenTableMangerRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *Handler) {
		rg.GET("/list", func(c *gin.Context) {
			handle.TableRoute.FindByPage(c, func(helper xgorm.IQueryHelper) *gorm.DB {
				return helper.SpliceQueryIfExit(c, map[string]xgorm.OperationType{})
			})
		})
		rg.GET("/:id", handle.TableRoute.FindOneById)
		rg.GET("/models", handle.GetTableList)
		rg.POST("/gen", handle.GenTables)
		rg.GET("/preview/:name", handle.GenPreView)
		rg.POST("", handle.TableRoute.Create)
		rg.PUT("/:id", handle.TableRoute.UpDateById)
		rg.DELETE("", handle.TableRoute.DeleteByIds)
	})
}
func regSysGenColumMangerRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *Handler) {
		rg.GET("/list", func(c *gin.Context) {
			handle.ColumRoute.FindByPage(c, func(helper xgorm.IQueryHelper) *gorm.DB {
				return helper.SpliceQueryIfExit(c, map[string]xgorm.OperationType{
					"tableName": xgorm.EQ,
				})
			})
		})
		rg.GET("/:name", handle.GetTableColumns)
		rg.PUT("/:id", handle.ColumRoute.UpDateById)
	})
}

type Handler struct {
	TableRoute xgorm.IRouterFunctions[model.SysGenTable]
	ColumRoute xgorm.IRouterFunctions[model.SysGenTableColumn]
	genService ISysGenService
}

func NewSysGenHandler(gen ISysGenService) *Handler {
	return &Handler{
		TableRoute: xgorm.InjectRouter[model.SysGenTable](gen.GetDB()),
		ColumRoute: xgorm.InjectRouter[model.SysGenTableColumn](gen.GetDB()),
		genService: gen,
	}
}

func (h Handler) GetTableList(c *gin.Context) {
	tablesInSql, err := h.genService.GetDB().Migrator().GetTables()
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	var tableNames []string
	for _, sqlTable := range tablesInSql {
		if !strings.HasPrefix(sqlTable, "sys_") {
			tableNames = append(tableNames, sqlTable)
		}
	}
	xresponse.SuccessCtx(c, tableNames)
	return
}
func (h Handler) GenTables(c *gin.Context) {
	var param GenTableParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err := h.genService.AddGenTable(c, param.Names)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
	return
}

func (h Handler) GetTableColumns(c *gin.Context) {
	var param GenTableColumnsParam
	if err := c.ShouldBindUri(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	columns, err := h.genService.GetTableColumns(c, param.Name)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, columns)
}

func (h Handler) GenPreView(c *gin.Context) {
	var param struct {
		TableName string `uri:"name" zh_comment:"表名" en_comment:"table name" validate:"required"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	preview, err := h.genService.CodePreview(c, param.TableName)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, preview)
}
