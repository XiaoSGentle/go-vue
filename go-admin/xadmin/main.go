package main

import (
	"xadmin/soybean"
	"xcore/core/xvariable"
)

func main() {
	soybean.GetAdminRouter().Run(xvariable.GlobalYmlConfig.GetString("HttpServer.Api.Port"))
}
