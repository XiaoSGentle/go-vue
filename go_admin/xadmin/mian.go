package main

import (
	"xcore/core/xadmin"
	"xcore/core/xvariable"
)

func main() {
	xadmin.GetSoybeanAdminRouter().Run(xvariable.GlobalYmlConfig.GetString("HttpServer.Api.Port"))
}
