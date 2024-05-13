package xmiddlewares

import (
	"github.com/gin-gonic/gin"
	"xcore/common/xvalidate"
	"xcore/core/xvariable"
)

// BindValidator
// Deprecated 目前默认是中文，如何修改还在想！
func BindValidator(c *gin.Context) {
	// 再做优化,目前先不用该中间件
	//header := c.GetHeader("Accept-Language")
	//if strings.HasPrefix(header, "zh") {
	//}
	//if strings.HasPrefix(header, "en") {
	//}
	c.Set(xvalidate.BIND_GIN_CONTEXT_KEY, xvariable.Validator)
}
