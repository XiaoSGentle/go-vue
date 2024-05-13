package soybean

import (
	inithandler "xadmin/soybean/init"
	"xadmin/soybean/system"
	"xcore/common/xauth"
	"xcore/common/xgorm"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

func GetAdminRouter() *xcore.GinCore {
	core := xcore.NewGinCore()
	core.RegisterRegFunction(xgorm.NewMysqlConnection)
	core.RegisterRegFunctions(system.Function)
	core.RegisterRouterGroups(system.Routers)

	core.AddRunBeforeFunctions([]func(){
		func() {
			xvariable.Auth = xauth.NewAuth(inithandler.GetRolesPermits())
		},
		func() {
			inithandler.InsertApisToSql(core.Router.Routes())
		},
	})
	return core
}
