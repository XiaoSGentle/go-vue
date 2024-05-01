package xvariable

import (
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"xcore/common/xcaptcha"
	"xcore/common/xconfig/interf"
)

var (
	// BasePath 项目运行路径
	BasePath string = "./"

	Logger *slog.Logger

	GlobalYmlConfig interf.YmlConfigInterf

	GormYmlConfig interf.YmlConfigInterf

	GormDB *gorm.DB

	Captcha *xcaptcha.Captcha
)
