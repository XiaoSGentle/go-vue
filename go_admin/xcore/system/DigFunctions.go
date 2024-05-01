package system

import (
	"xcore/core/xcore"
	"xcore/system/sys_auth"
	"xcore/system/sys_menu"
	"xcore/system/sys_role"
	"xcore/system/sys_user"

	"xcore/system/sys_route"
)

type System struct {
}

var (
	Routers = []*xcore.GroupBase{
		sys_auth.AuthGroup,
		sys_route.RouteGroup,
		sys_route.NoAuthRouteGroup,
		sys_route.SysMangerRoute,
		sys_menu.SysMenuGroup,
		sys_user.SysUserGroup,
		sys_role.SysRoleGroup,
	}
	Function = []interface{}{
		sys_auth.NewAuthService,
		sys_route.NewRouteService,
		sys_user.NewSysUserService,
		sys_menu.NewSysMenuService,
		sys_role.NewSysRoleService,
	}
)
