package sys_role

import (
	"github.com/gin-gonic/gin"
	inithandler "xadmin/soybean/init"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	baseType "xcore/common/xtype/xbase"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var SysRoleGroup = xcore.Group("/system/role", newSysRoleHandler, regSysRole, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize, xmiddlewares.Verify)

func regSysRole(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysRoleHandler) {
		rg.GET("/list", handle.SysRoleList)
		rg.GET("/:code", handle.GetRolePermit)
		rg.POST("", handle.AddSysRole)
		rg.PUT("/:id", handle.UpdateRole)
		rg.DELETE("", handle.DeleteByIds)
		rg.PUT("/apis", handle.UpdateApis)
		rg.PUT("/menus", handle.UpdateMenus)
		rg.PUT("/home", handle.UpdateHome)
	})
}

type sysRoleHandler struct {
	service ISysRoleService
}

func newSysRoleHandler(auth ISysRoleService) *sysRoleHandler {
	return &sysRoleHandler{service: auth}
}

func (r sysRoleHandler) SysRoleList(c *gin.Context) {
	var sysUserListParam SysRoleListParam
	if err := c.ShouldBind(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&sysUserListParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	list, err := r.service.SysRoleList(c, &sysUserListParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)
}

func (r sysRoleHandler) AddSysRole(c *gin.Context) {
	var param AddOrUpdateSysRoleParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	err := r.service.AddRole(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}

func (r sysRoleHandler) UpdateRole(c *gin.Context) {
	var param AddOrUpdateSysRoleParam
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
	err = r.service.UpdateRole(c, pathId.Id, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	xresponse.UpdateSuccessCtx(c)
}

func (r sysRoleHandler) DeleteByIds(c *gin.Context) {
	var ids baseType.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := r.service.DeleteRole(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}

func (r sysRoleHandler) GetRolePermit(c *gin.Context) {
	var roleCode struct {
		RoleCode string `uri:"code"`
	}
	err := c.BindUri(&roleCode)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	resp, err := r.service.GetRolePermit(c, roleCode.RoleCode)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, resp)
}

func (r sysRoleHandler) UpdateApis(c *gin.Context) {
	var param UpdateRolePermitParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err := r.service.UpdateApis(c, param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xvariable.Auth.RefreshAuthStore(inithandler.GetRolesPermits())
	xresponse.UpdateSuccessCtx(c)
}

func (r sysRoleHandler) UpdateMenus(c *gin.Context) {
	var param UpdateRolePermitParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err := r.service.UpdateMenus(c, param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xvariable.Auth.RefreshAuthStore(inithandler.GetRolesPermits())
	xresponse.UpdateSuccessCtx(c)
}

func (r sysRoleHandler) UpdateHome(c *gin.Context) {
	var param UpdateRoleHomeParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvariable.Validator.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err := r.service.UpdateHome(c, param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.UpdateSuccessCtx(c)
}
