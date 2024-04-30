package xgorm

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
	"xcore/core/xvariable"
)

func GetMysqlConnection() *gorm.DB {
	Host := xvariable.GormYmlConfig.GetString("Gorm.Mysql.Host")
	DataBase := xvariable.GormYmlConfig.GetString("Gorm.Mysql.DataBase")
	Port := xvariable.GormYmlConfig.GetInt64("Gorm.Mysql.Port")
	User := xvariable.GormYmlConfig.GetString("Gorm.Mysql.User")
	Pass := xvariable.GormYmlConfig.GetString("Gorm.Mysql.Pass")
	Charset := xvariable.GormYmlConfig.GetString("Gorm.Mysql.Charset")
	// ?
	SetMaxIdleConn := xvariable.GormYmlConfig.GetInt("Gorm.Mysql.SetMaxIdleConn")
	// 最大连接数
	SetMaxOpenConn := xvariable.GormYmlConfig.GetInt("Gorm.Mysql.SetMaxOpenConn")
	// ?
	SetConnMaxLifetime := xvariable.GormYmlConfig.GetDuration("Gorm.Mysql.SetConnMaxLifetime")
	// sql执行时间超过此时间单位（秒），就会触发系统日志记录
	//SlowThreshold := xvariable.GormYmlConfig.GetInt("Gorm.Mysql.SlowThreshold")

	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", User, Pass, Host, Port, DataBase, Charset)

	mysqlDialectic := mysql.Open(connectUrl)

	gormDb, err := gorm.Open(mysqlDialectic, &gorm.Config{})
	if err != nil {
		xvariable.Logger.Error(context.Background(), "gorm 初始化出错:"+err.Error())
		// >
		panic(err.Error())
	}
	resolverConf := dbresolver.Config{
		Replicas: []gorm.Dialector{mysqlDialectic}, //  读 操作库，查询类
		Policy:   dbresolver.RandomPolicy{},        // sources/replicas 负载均衡策略适用于
	}
	err = gormDb.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).SetConnMaxLifetime(SetConnMaxLifetime * time.Second).SetMaxIdleConns(SetMaxIdleConn).SetMaxOpenConns(SetMaxOpenConn))
	if err != nil {
		xvariable.Logger.Error(context.Background(), "gorm dbResolver 出错:"+err.Error())
	}
	return gormDb
}
