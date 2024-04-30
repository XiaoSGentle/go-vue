package xmiddlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"xcore/common/xerror"
	"xcore/common/xresponse"
	xtoken "xcore/common/xtoken/jwt"
	"xcore/core/xvariable"
)

func Authorize(c *gin.Context) {
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	//jwtTokenCreatedExpireAt := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")
	claims, err := xtoken.GetClaimsByRequest(c, jwtTokenSignKey)
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			if time.Now().Before(claims.ExpiresAt.Add(time.Duration(jwtTokenRefreshAllowSec) * time.Second)) {
				xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.TOKEN_EXPIRE_ERROR))
				c.Abort()
			}
		} else {
			xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.TOKEN_EXPIRE_ERROR))
			c.Abort()
		}
	}

	c.Next()
}
