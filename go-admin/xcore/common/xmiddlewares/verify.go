package xmiddlewares

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xerror"
	"xcore/common/xresponse"
	"xcore/common/xtoken"
	"xcore/core/xvariable"
)

// Verify 是否有权限
func Verify(c *gin.Context) {
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	var payLoad, err = xtoken.GetPayloadByRequest(c, jwtTokenSignKey)
	if err != nil {
		xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.TOKEN_ERROR))
		c.Abort()
		return
	}
	if !xvariable.Auth.CheckApiPermit(payLoad.Roles, c.Request.Method, c.FullPath()) {
		xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.NO_RERMIT_ERROR))
		c.Abort()
		return
	}
	c.Next()
}
