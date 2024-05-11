package xmiddlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"xcore/common/xerror"
	"xcore/common/xresponse"
	"xcore/common/xtoken"
	"xcore/core/xvariable"
)

// Authorize 验证Token是否过期，以及有效
func Authorize(c *gin.Context) {
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")
	claims, err := xtoken.GetClaimsByRequest(c, jwtTokenSignKey)
	c.Set(xtoken.ContextKey, claims.ClaimsPayload)
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			if time.Now().Before(claims.ExpiresAt.Add(time.Duration(jwtTokenRefreshAllowSec) * time.Second)) {
				xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.TOKEN_EXPIRE_ERROR))
				c.Abort()
			}
		} else {
			xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.TOKEN_ERROR))
			c.Abort()
		}
	}
	c.Next()
}
