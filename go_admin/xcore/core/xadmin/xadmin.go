package xadmin

import (
	"xcore/common/xgorm"
	"xcore/core/xcore"
	"xcore/system"
)

func GetSoybeanAdminRouter() *xcore.GinCore {
	core := xcore.NewGinCore()
	core.RegisterRegFunction(xgorm.GetMysqlConnection)
	core.RegisterRegFunctions(system.Function)
	core.RegisterRouterGroups(system.Routers)
	return core
}
