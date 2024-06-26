package main

import (
	"xadmin/soybean"
	"xcore/core/xvariable"
)

func main() {
	soybean.GetAdminRouter().Run(xvariable.GlobalYmlConfig.GetString("HttpServer.Api.Port"))
}

//func main() {
//	var data = xencrypt.Base64Md5("Xiaos123!")
//	println(data)
//}
