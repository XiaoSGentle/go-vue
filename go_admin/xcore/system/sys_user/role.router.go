package sys_user

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
)

var SysUserGroup = xcore.Group("/systemManage/user", newSysUserHandler, regSysUser, xmiddlewares.LogMiddleHandler)

func regSysUser(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysUserHandler) {
		rg.GET("/list", handle.SysUserList)
	})
}

type sysUserHandler struct {
	authService ISysUserService
}

func newSysUserHandler(auth ISysUserService) *sysUserHandler {
	return &sysUserHandler{authService: auth}
}

func (a sysUserHandler) SysUserList(c *gin.Context) {
	var sysUserListParam SysUserListParam
	if err := c.ShouldBind(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return

	}
	list, err := a.authService.SysUserList(c, &sysUserListParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)
}
