package sys_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	xtoken "xcore/common/xtoken/jwt"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var AuthGroup = xcore.Group("/auth", newAuthHandler, regAuth, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)
var NoAuthAuthGroup = xcore.Group("/auth", newAuthHandler, regNoAuthAuth, xmiddlewares.LogMiddleHandler)

func regAuth(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *authHandler) {
		rg.GET("/getUserInfo", handle.GetUserInfo)
	})
}
func regNoAuthAuth(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *authHandler) {
		rg.POST("/login", handle.UserLogin)
		rg.POST("/refreshToken", handle.RefreshToken)
	})
}

type authHandler struct {
	authService IAuthService
}

func newAuthHandler(auth IAuthService) *authHandler {
	return &authHandler{authService: auth}
}

func (a authHandler) GetUserInfo(c *gin.Context) {
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	payload, err := xtoken.GetPayloadByRequest(c, jwtTokenSignKey)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	userInfInSql, err := a.authService.GetUserById(c, payload.Uid)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, UserInfoVo{
		UserId:   fmt.Sprintf("%d", userInfInSql.ID),
		UserName: userInfInSql.Username,
		Roles:    strings.Split(userInfInSql.Roles, ","),
		Apis:     xvariable.Auth.GetApiCodesByRoles(strings.Split(userInfInSql.Roles, ",")),
	})
}
func (a authHandler) RefreshToken(c *gin.Context) {
	var refreshTokenParam RefreshTokenParam
	if err := c.ShouldBind(&refreshTokenParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(refreshTokenParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	jwtTokenRefreshSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenRefreshSignKey")
	jwtTokenCreatedExpireAt := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")
	payload, err := xtoken.GetPayloadByRequest(c, jwtTokenSignKey)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	userInfInSql, err := a.authService.GetUserById(c, payload.Uid)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	jwtToken, refreshToken, err := xtoken.RefreshToken(
		refreshTokenParam.RefreshToken, jwtTokenRefreshSignKey,
		jwtTokenRefreshAllowSec, jwtTokenSignKey, jwtTokenCreatedExpireAt,
		&xtoken.ClaimsPayload{
			Uid:      userInfInSql.ID,
			NickName: userInfInSql.Nickname,
			Roles:    strings.Split(userInfInSql.Roles, ",")})
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	xresponse.SuccessCtx(c, LoginVo{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	})

}
func (a authHandler) UserLogin(c *gin.Context) {
	var loginParam LoginParam
	if err := c.ShouldBind(&loginParam); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(loginParam); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	LoginVo, err := a.authService.UserLogin(c, &loginParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, LoginVo)
}
