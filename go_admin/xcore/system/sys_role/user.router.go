package sys_role

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
)

var SysRoleGroup = xcore.Group("/systemManage/role", newSysRoleHandler, regSysRole, xmiddlewares.LogMiddleHandler)

func regSysRole(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysRoleHandler) {
		rg.GET("/list", handle.SysRoleList)
	})
}

type sysRoleHandler struct {
	authService ISysRoleService
}

func newSysRoleHandler(auth ISysRoleService) *sysRoleHandler {
	return &sysRoleHandler{authService: auth}
}

func (a sysRoleHandler) SysRoleList(c *gin.Context) {
	var sysUserListParam SysRoleListParam
	if err := c.ShouldBind(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	list, err := a.authService.SysRoleList(c, &sysUserListParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)
}
