package system

import (
	"xadmin/soybean/system/sys_auth"
	"xadmin/soybean/system/sys_cron"
	"xadmin/soybean/system/sys_department"
	"xadmin/soybean/system/sys_dict"
	"xadmin/soybean/system/sys_log"
	"xadmin/soybean/system/sys_menu"
	"xadmin/soybean/system/sys_role"
	"xadmin/soybean/system/sys_route"
	"xadmin/soybean/system/sys_user"
	"xcore/core/xcore"
)

type System struct {
}

var (
	RouterGroups = []*xcore.GroupBase{
		sys_auth.AuthGroup,
		sys_route.RouteGroup,
		sys_route.NoAuthRouteGroup,
		sys_auth.NoAuthAuthGroup,
		sys_route.SysMangerRoute,
		sys_dict.NoAuthDictGroup,
		sys_dict.SysDictGroup,
		sys_dict.SysDictDataGroup,
		sys_menu.SysMenuGroup,
		sys_user.SysUserGroup,
		sys_role.SysRoleGroup,
		sys_log.LogGroup,
		sys_log.NoLogLogGroup,
		sys_department.SysDepartmentGroup,
		sys_cron.SysCronGroup,
	}
	NewServiceFunctions = []interface{}{
		sys_auth.NewAuthService,
		sys_route.NewRouteService,
		sys_user.NewSysUserService,
		sys_menu.NewSysMenuService,
		sys_dict.NewSysDictService,
		sys_dict.NewSysDictDataService,
		sys_role.NewSysRoleService,
	}
)
