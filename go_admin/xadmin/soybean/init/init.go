package init_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
	"strings"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xauth"
	"xcore/core/xvariable"
)

func InsertApisToSql(routers []gin.RouteInfo) {
	xvariable.Logger.InfoLog.Info("insert apis to sql start...")
	println("insert apis to sql start...")
	apiQuery := query.Use(xvariable.GormDB).SysAPI

	// 清除所欲路由
	_, _ = apiQuery.Where(apiQuery.APICode.IsNotNull()).Delete()

	for _, router := range routers {
		if !strings.HasPrefix(router.Path, "/debug/") {
			_ = apiQuery.Create(&model.SysAPI{
				APICode:       fmt.Sprintf("%s::%s", router.Method, router.Path),
				Version:       0,
				SoftDeleteTag: 0,
				UpdateTime:    time.Now(),
				UpdateUID:     0,
				CreateUID:     0,
				CreateBy:      "",
				CreateTime:    time.Now(),
				UpdateBy:      "",
			})

		}

	}
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
