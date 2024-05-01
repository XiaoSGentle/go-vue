package sys_menu

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/core/xcore"
)

var SysMenuGroup = xcore.Group("/systemManage", newSysMenuHandler, regSysMenu, xmiddlewares.LogMiddleHandler)

func regSysMenu(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysMenuHandler) {
		rg.GET("/getMenuList", handle.GetMenuList)
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
