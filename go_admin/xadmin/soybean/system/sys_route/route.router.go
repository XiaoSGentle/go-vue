package sys_route

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/common/xslice"
	xtoken "xcore/common/xtoken/jwt"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var RouteGroup = xcore.Group("/route", newRouteHandler, regRoute, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)
var NoAuthRouteGroup = xcore.Group("/route", newRouteHandler, regNoAuthRoute, xmiddlewares.LogMiddleHandler)
var SysMangerRoute = xcore.Group("/system/route", newRouteHandler, regSysMangerRoute, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize, xmiddlewares.Verify)

func regSysMangerRoute(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *routeHandler) {
		rg.GET("/pages", handle.GetAllPages)
		rg.GET("/apis", handle.GetAllApis)
		rg.GET("/tree", handle.GetMenuTree)
		rg.GET("/roles", handle.GetAllRoles)
	})
}

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
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	payload, _ := xtoken.GetPayloadByRequest(c, jwtTokenSignKey)
	routers, err := a.authService.GetUserRouters(c, xslice.StringToInt32(xvariable.Auth.GetMenuIdsByRoles(payload.Roles)))
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
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
func (a routeHandler) GetAllPages(c *gin.Context) {
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	customClaim, err := xtoken.GetPayloadByRequest(c, jwtTokenSignKey)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	pages, err := a.authService.GetAllPages(c, customClaim.Roles)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, pages)
}

func (a routeHandler) IsRouteExist(c *gin.Context) {

}

func (a routeHandler) GetMenuTree(c *gin.Context) {
	pages, err := a.authService.GetMenuTreeSimple(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, pages)
}

func (a routeHandler) GetAllApis(c *gin.Context) {
	apis, err := a.authService.GetALLApis(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, apis)
	return
}

func (a routeHandler) GetAllRoles(c *gin.Context) {
	apis, err := a.authService.GetAllRoles(c)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, apis)
}
