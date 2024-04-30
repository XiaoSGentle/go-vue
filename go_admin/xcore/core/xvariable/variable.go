package xvariable

import (
	"gorm.io/gorm"
	"xcore/common/xconfig/interf"
	logger2 "xcore/common/xlogger"
)

var (
	// BasePath 项目运行路径
	BasePath string = "./"

	Logger *logger2.Logger

	GlobalYmlConfig interf.YmlConfigInterf

	GormYmlConfig interf.YmlConfigInterf

	GormDB *gorm.DB
)
