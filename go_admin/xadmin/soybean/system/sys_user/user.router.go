package sys_user

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	baseType "xcore/common/xtype/xbase"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var SysUserGroup = xcore.Group("/system/user", newSysUserHandler, regSysUser, xmiddlewares.LogMiddleHandler)

func regSysUser(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysUserHandler) {
		rg.GET("/list", handle.SysUserList)
		rg.POST("", handle.AddSysUser)
		rg.PUT("/:id", handle.UpdateUser)
		rg.DELETE("", handle.DeleteByIds)
	})
}

type sysUserHandler struct {
	service ISysUserService
}

func newSysUserHandler(auth ISysUserService) *sysUserHandler {
	return &sysUserHandler{service: auth}
}

func (h sysUserHandler) SysUserList(c *gin.Context) {
	var sysUserListParam SysUserListParam
	if err := c.ShouldBind(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return

	}
	list, err := h.service.SysUserList(c, &sysUserListParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)
}

func (h sysUserHandler) AddSysUser(c *gin.Context) {
	var param AddOrUpdateSysUserParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	err := h.service.AddUser(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}

func (h sysUserHandler) UpdateUser(c *gin.Context) {
	var param AddOrUpdateSysUserParam
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
	err = h.service.UpdateUser(c, pathId.Id, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.UpdateSuccessCtx(c)
}

func (h sysUserHandler) DeleteByIds(c *gin.Context) {
	var ids baseType.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvariable.Validator.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := h.service.DeleteUser(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}
