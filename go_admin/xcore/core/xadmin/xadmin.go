package xadmin

import (
	"xcore/common/xgorm"
	"xcore/core/xcore"
	"xcore/system"
)

func GetSoybeanAdminRouter() *xcore.GinCore {
	core := xcore.NewGinCore()
	core.RegisterOneRegFunction(xgorm.GetMysqlConnection)
	core.RegisterRouterGroupArray(system.Routers)
	core.RegisterRegFunctionArray(system.Function)
	return core
}
