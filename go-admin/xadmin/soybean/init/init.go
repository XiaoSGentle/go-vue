package init_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
	"strings"
	"time"
	"xadmin/soybean/crons"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xauth"
	"xcore/common/xcron"
	"xcore/common/xtype/xslice"
	"xcore/core/xconst"
	"xcore/core/xvariable"
)

func InsertApisToSql(routers []gin.RouteInfo) {
	xvariable.Logger.InfoLog.Info("insert apis to sql...")
	println("insert apis to sql...")
	apiQuery := query.Use(xvariable.GormDB).SysAPI
	// 清除所有路由
	_, _ = apiQuery.Where(apiQuery.APICode.IsNotNull()).Delete()
	var insertModels []*model.SysAPI
	for _, router := range routers {
		if !strings.HasPrefix(router.Path, "/debug/") {
			insertModels = append(insertModels, &model.SysAPI{
				APICode:    fmt.Sprintf("%s::%s", router.Method, router.Path),
				UpdateTime: time.Now(),
				UpdateUID:  0,
				CreateUID:  0,
				CreateBy:   "",
				CreateTime: time.Now(),
				UpdateBy:   "",
			})
		}
	}
	_ = apiQuery.Create(insertModels...)
	xvariable.Logger.InfoLog.Info("insert apis to sql success!")
	println("insert apis to sql success!")

}

func GetRolesPermits() []xauth.AuthContent {
	xvariable.Logger.InfoLog.Info("start InitRolesPermits...")
	println("start InitRolesPermits...")
	var roles []xauth.AuthContent
	roleQuery := query.Use(xvariable.GormDB).SysRole
	rolesInfos, err := roleQuery.Find()
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
	}
	for _, roleInfo := range rolesInfos {
		roles = append(roles, xauth.AuthContent{
			RoleCode: roleInfo.Code,
			Menus:    strings.Split(roleInfo.MenuIds, ","),
			Apis:     strings.Split(roleInfo.APICodes, ","),
		})
	}
	xvariable.Logger.InfoLog.Info("InitRolesPermits success!")
	println("InitRolesPermits success!")
	return roles
}

func InitCron() xcron.IXCron {
	println("init Cron...")
	xCorn := xcron.NewXCorn(crons.AllCron{})
	cronQuery := query.Use(xvariable.GormDB).SysCron
	// 扫描一个结构体下所有的函数
	keys, functions := xCorn.ScannerFunctions()
	// 先删除数据库中有但是代码中不存在的
	_, _ = cronQuery.WithContext(context.Background()).Where(cronQuery.Key.NotIn(keys...)).Delete()
	find, _ := cronQuery.Select(cronQuery.Key).Find()
	keysInSql := extractKeys(find)
	// 写入数据库
	for _, function := range functions {
		if !xslice.StringExist(keysInSql, function.Key) {
			_ = cronQuery.Create(&model.SysCron{
				Key:           function.Key,
				Description:   function.Key,
				Schedule:      "",
				Status:        xconst.StatusBanned,
				Arguments:     strings.Join(function.ParamTypes, ","),
				ArgumentsType: strings.Join(function.ParamTypes, ","),
			})
		}
	}
	// 设置从数据库添加定时任务的数据
	xCorn.SetCronRule(func(cron *cron.Cron) {
		cronInSql, _ := cronQuery.Find()
		for _, sysCron := range cronInSql {
			if sysCron.Status == xconst.StatusOK {
				_ = cron.AddFunc(sysCron.Schedule, func() {
					var interfaceSlice []interface{}
					for _, arg := range strings.Split(sysCron.Arguments, ",") {
						interfaceSlice = append(interfaceSlice, arg)
					}
					xCorn.CallMethodsOnValue(sysCron.Key, interfaceSlice...)
				})
			}
		}
	})
	// 启动定时任务
	xCorn.StartCorn()
	println("init Cron Success!")
	return xCorn
}

func extractKeys(results []*model.SysCron) []string {
	keys := make([]string, len(results))
	for i, result := range results {
		keys[i] = result.Key
	}
	return keys
}
