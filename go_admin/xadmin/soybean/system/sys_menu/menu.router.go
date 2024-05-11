package sys_menu

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	basetype "xcore/common/xtype/xbase"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
)

var SysMenuGroup = xcore.Group("/system/menu", newSysMenuHandler, regSysMenu, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize, xmiddlewares.Verify)

func regSysMenu(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysMenuHandler) {
		rg.GET("/list", handle.GetMenuList)
		rg.POST("", handle.AddMenu)
		rg.PUT("/:id", handle.UpdateMenu)
		rg.DELETE("", handle.DeleteMenuByIds)
	})
}

type sysMenuHandler struct {
	authService ISysMenuService
}

func newSysMenuHandler(auth ISysMenuService) *sysMenuHandler {
	return &sysMenuHandler{authService: auth}
}

func (h sysMenuHandler) GetMenuList(c *gin.Context) {
	list, err := h.authService.GetMenuList(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)

}

func (h sysMenuHandler) AddMenu(c *gin.Context) {
	var sysMenuAddParam AddOrUpDateSysMenuParam
	if err := c.ShouldBind(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	err := h.authService.AddMenu(c, &sysMenuAddParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}

func (h sysMenuHandler) UpdateMenu(c *gin.Context) {
	var sysMenuAddParam AddOrUpDateSysMenuParam
	if err := c.ShouldBind(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	var pathId basetype.DetailsId
	err := c.BindUri(&pathId)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err = h.authService.UpdateMenu(c, pathId.Id, &sysMenuAddParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.UpdateSuccessCtx(c)
}

func (h sysMenuHandler) DeleteMenuByIds(c *gin.Context) {
	var ids basetype.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := h.authService.DeleteMenu(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}
