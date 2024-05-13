package xvariable

import (
	"gorm.io/gorm"
	"xcore/common/xauth"
	"xcore/common/xcaptcha"
	"xcore/common/xconfig/interf"
	"xcore/common/xcron"
	"xcore/common/xlogger"
	"xcore/common/xvalidate"
)

var (
	// BasePath 项目运行路径
	BasePath string = "./"

	Logger *xlogger.Logger

	GlobalYmlConfig interf.YmlConfigInterf

	GormYmlConfig interf.YmlConfigInterf

	GormDB *gorm.DB

	Captcha *xcaptcha.Captcha

	Auth *xauth.Auth

	Validator xvalidate.IValidator

	XCron xcron.IXCron
)
