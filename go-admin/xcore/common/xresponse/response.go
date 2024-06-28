package xresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"xcore/common/xerror"
	"xcore/core/xvariable"
)

func SuccessCtx(c *gin.Context, v any) any {
	success := createSuccess(v)
	ReturnJson(c, 200, success)
	c.Abort()
	return nil
}
func CreateSuccessCtx(c *gin.Context) {
	success := createSuccess("创建成功")
	ReturnJson(c, 200, success)
	c.Abort()
}

func UpdateSuccessCtx(c *gin.Context) {
	success := createSuccess("更新成功")
	ReturnJson(c, 200, success)
	c.Abort()
}

func DeleteSuccessCtx(c *gin.Context) {
	success := createSuccess("删除成功")
	ReturnJson(c, 200, success)
	c.Abort()
}

func ErrorCtx(c *gin.Context, err error) any {
	errCode := xerror.SERVER_COMMON_ERROR
	errMsg := "服务器开小差啦，稍后再来试一试"
	causeErr := errors.Cause(err)
	var codeError *xerror.CodeError
	// 自定义错误 比如影响行数为0等等
	if errors.As(causeErr, &codeError) {
		errCode = codeError.GetErrCode()
		errMsg = codeError.GetErrMsg()
	} else {
		// 不是的话就是系统错误
		// todo: 这里有条件再进行优化
		// 写入日志
		xvariable.Logger.ErrorLog.ErrorContext(c, "响应失败错错误信息"+err.Error())
	}
	ReturnJson(c, http.StatusOK, createError(errCode, errMsg))
	c.Abort()
	return nil
}

func createSuccess(v any) *ResponseSuccess {
	if v == nil {
		return &ResponseSuccess{
			Code: 200,
			Msg:  "SUCCESS",
		}
	}
	return &ResponseSuccess{
		Code: 200,
		Msg:  "SUCCESS",
		Data: v,
	}
}

func createError(code uint32, msg string) *ResponseError {
	return &ResponseError{
		Code: code,
		Msg:  msg,
	}
}

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func ReturnJson(Context *gin.Context, httpCode int, data interface{}) {
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, data)
}
