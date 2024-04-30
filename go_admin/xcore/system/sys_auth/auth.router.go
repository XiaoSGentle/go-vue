package sys_auth

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	xtoken "xcore/common/xtoken/jwt"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
	"xcore/core/xvariable"
)

var AuthGroup = xcore.Group("/auth", newAuthHandler, regMenu, xmiddlewares.LogMiddleHandler)

func regMenu(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *authHandler) {
		rg.POST("/login", handle.UserLogin)
		rg.POST("/refreshToken", handle.RefreshToken)
		rg.GET("/getUserInfo", handle.GetUserInfo)
		rg.GET("/oauth2/callback", handle.Oauth2CallBack)
		rg.POST("/oauth2/bind", handle.Oauth2Bind)
		rg.GET("/captcha_data", handle.GetCapture)
		rg.GET("/maxkey_redirect", handle.MaxKeyRedirect)
		rg.POST("/captcha_check", handle.CheckCapture)
		rg.POST("/get_sms", handle.GetSms)
		rg.POST("/sms_login", handle.SmsCodeLogin)
	})
}

type authHandler struct {
	authService IAuthService
}

func newAuthHandler(auth IAuthService) *authHandler {
	return &authHandler{authService: auth}
}
func (a authHandler) Oauth2Bind(c *gin.Context) {

}

func (a authHandler) MaxKeyRedirect(c *gin.Context) {
}

// Oauth2CallBack Oauth2登录认证
func (a authHandler) Oauth2CallBack(c *gin.Context) {
	//// 获取一些变量
	//clientId := variable.ConfigYml.GetString("Oauth2.MaxKey.ClientId")
	//clientSecret := variable.ConfigYml.GetString("Oauth2.MaxKey.ClientSecret")
	//redirectUri := variable.ConfigYml.GetString("Oauth2.MaxKey.RedirectUri")
	//maxKeyBaseUrl := variable.ConfigYml.GetString("Oauth2.MaxKey.AuthUri")
	//// 获取code
	//callBackCode := c.Query("code")
	//// 授权类型
	//grantType := "authorization_code"
	//
	//// 获取前端的地址
	//fontBaseUrl := variable.ConfigYml.GetString("Oauth2.FrontBaseUrl")
	//// 获取assess Token Url
	//accessTokenUrl := fmt.Sprintf(`%s/sign/authz/oauth/v20/token?client_id=%s&client_secret=%s&grant_type=%s&redirect_uri=%s&code=%s`,
	//	maxKeyBaseUrl, clientId, clientSecret, grantType, redirectUri, callBackCode)
	////发起请求
	//accessTokenResp, err := http.Get(accessTokenUrl)
	//if err != nil {
	//	return
	//}
	//// 获取消息体并转换为Json
	//defer accessTokenResp.Body.Close()
	//body, err := io.ReadAll(accessTokenResp.Body)
	//fmt.Println("原始值 accessTokenResp = " + string(body))
	//var accessToken Oauth2MaxKeyAccessToken
	//if err := json.Unmarshal(body, &accessToken); err != nil {
	//	fmt.Println("转换accesstoken 出错" + err.Error())
	//}
	//fmt.Println("accesstoken值为：")
	//// 发起请求并处理
	//fmt.Println(accessToken)
	//
	//// 拼接获取用户的信息
	//getUserInfoUrl := fmt.Sprintf("%s/sign/api/oauth/v20/me?access_token=%s", maxKeyBaseUrl, accessToken.AccessToken)
	//// 请求并转化用户信息
	//userInfoResp, err := http.Get(getUserInfoUrl)
	//defer userInfoResp.Body.Close()
	//userInfoByte, err := io.ReadAll(userInfoResp.Body)
	//fmt.Println("原始值 userInfoByte = " + string(userInfoByte))
	//var userInfo Oauth2MaxKeyUserInfo
	//if err := json.Unmarshal(userInfoByte, &userInfo); err != nil {
	//	fmt.Println("json.Unmarshal 出错")
	//	fmt.Println("err" + err.Error())
	//	response.Fail(c, consts.Oauth2LoginFailCode, consts.Oauth2LoginFailMsg, false)
	//	return
	//}
	//fmt.Println("过期时间")
	//fmt.Println(accessToken.ExpiresIn)
	////
	//authToken, states, bindKey, err := a.authService.Oauth2Login(userInfo.UserId, consts.Oauth2ServiceMaxKey, int64(accessToken.ExpiresIn))
	//if err != nil {
	//	response.Fail(c, consts.Oauth2LoginFailCode, consts.Oauth2LoginFailMsg, false)
	//	return
	//} else if strings.Compare(states, consts.Oauth2ServiceNotBind) == 0 {
	//	// 未绑定跳转未绑定
	//	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/#/login/bind-reacool-center?bindKey=%s", fontBaseUrl, bindKey))
	//	return
	//}
	//// 跳转登录页面
	//c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/#/login?token=%s&refreshToken=%s", fontBaseUrl, authToken.Token, authToken.RefreshToken))
	//return
}
func (a authHandler) GetSms(c *gin.Context) {
	println("get sms")
}
func (a authHandler) SmsCodeLogin(c *gin.Context) {

}
func (a authHandler) GetUserInfo(c *gin.Context) {
	xresponse.SuccessCtx(c, UserInfoVo{
		UserId:   "10001",
		UserName: "Xiaos",
		Roles:    []string{"R_SUPER"},
		Buttons:  []string{""},
	})
}
func (a authHandler) RefreshToken(c *gin.Context) {
	var refreshTokenParam RefreshTokenParam
	if err := c.ShouldBind(&refreshTokenParam); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(refreshTokenParam); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	jwtTokenCreatedExpireAt := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")

	jwtToken, refreshToken, err := xtoken.RefreshToken(refreshTokenParam.RefreshToken, jwtTokenSignKey, jwtTokenRefreshAllowSec, jwtTokenSignKey, jwtTokenCreatedExpireAt, 1, []string{
		"R_SUPER",
	})
	if err != nil {
		xresponse.ErrorCtx(c, err)
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
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	jwtTokenCreatedExpireAt := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")
	jwtToken, err := xtoken.GenerateJwtToken(jwtTokenSignKey, jwtTokenCreatedExpireAt, 1, []string{"R_SUPER"})
	refreshToken, err := xtoken.GenerateJwtToken(jwtTokenSignKey, jwtTokenRefreshAllowSec, 1, []string{"R_SUPER"})
	if err != nil {
		xresponse.ErrorCtx(c, err)
	}
	xresponse.SuccessCtx(c, LoginVo{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	})
}
func (a authHandler) GetCapture(c *gin.Context) {

}
func (a authHandler) CheckCapture(c *gin.Context) {

}
