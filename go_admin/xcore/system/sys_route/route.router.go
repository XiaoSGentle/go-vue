package sys_route

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/core/xcore"
)

var RouteGroup = xcore.Group("/route", newRouteHandler, regRoute, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)
var NoAuthRouteGroup = xcore.Group("/route", newRouteHandler, regNoAuthRoute, xmiddlewares.LogMiddleHandler)

func regRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *routeHandler) {
		rg.GET("/getUserRoutes", handle.GetUserRoutes)
		rg.GET("/isRouteExist", handle.IsRouteExist)
	})
}

func regNoAuthRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *routeHandler) {
		rg.GET("/getConstantRoutes", handle.GetConstantRoutes)
	})
}

type routeHandler struct {
	authService IRouteService
}

func newRouteHandler(auth IRouteService) *routeHandler {
	return &routeHandler{authService: auth}
}

func (a routeHandler) GetUserRoutes(c *gin.Context) {
	routers, err := a.authService.GetUserRouters(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
	}
	xresponse.SuccessCtx(c, routers)
}
func (a routeHandler) GetConstantRoutes(c *gin.Context) {
	routers, err := a.authService.GetConstantRoutes(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
	}
	xresponse.SuccessCtx(c, routers)

}
func (a routeHandler) IsRouteExist(c *gin.Context) {

}
