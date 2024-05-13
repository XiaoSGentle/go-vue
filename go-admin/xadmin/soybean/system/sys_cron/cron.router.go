package sys_cron

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xadmin/soybean/dao/model"
	"xcore/common/xgorm"
	"xcore/common/xresponse"
	baseType "xcore/common/xtype/xbase"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var SysCronGroup = xcore.Group("/system/cron", NewCronHandler, regCronMangerRoute)

func regCronMangerRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *Handler) {
		rg.GET("/list", handle.FindCronListByPage)
		rg.PUT("", handle.UpdateCron)
	})
}

type Handler struct {
	xgorm.IRouterFunctions[model.SysCron]
	service ISysCronService
}

func NewCronHandler(db *gorm.DB) *Handler {
	return &Handler{
		IRouterFunctions: xgorm.InjectRouter[model.SysCron](db),
		service:          NewSysDepartService(db),
	}
}

func (h Handler) FindCronListByPage(c *gin.Context) {
	var param baseType.PageParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return

	}
	list, err := h.service.FindCronList(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)
}
func (h Handler) UpdateCron(c *gin.Context) {
	var param UpdateCronReq
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	var pathId baseType.DetailsId
	err := c.BindUri(&pathId)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err = h.service.UpdateCron(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xvariable.XCron.ReStartCorn()
	xresponse.UpdateSuccessCtx(c)
}
