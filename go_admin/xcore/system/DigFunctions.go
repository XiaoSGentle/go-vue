package system

import (
	"xcore/core/xcore"
	"xcore/system/sys_auth"
	"xcore/system/sys_route"
)

type System struct {
}

var (
	Routers = []*xcore.GroupBase{
		sys_auth.AuthGroup,
		sys_route.RouteGroup,
		sys_route.NoAuthRouteGroup,
	}
	Function = []interface{}{
		sys_auth.NewAuthService,
		sys_route.NewRouteService,
	} // Routers
)
