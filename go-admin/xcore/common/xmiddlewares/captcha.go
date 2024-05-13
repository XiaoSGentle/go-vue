package xmiddlewares

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xresponse"
	"xcore/core/xvariable"
)

func CheckCaptcha(c *gin.Context) {
	ok, err := xvariable.Captcha.VerifyCaptcha(c)
	if err != nil || !ok {
		xresponse.ErrorCtx(c, err)
	}
	c.Next()
}
