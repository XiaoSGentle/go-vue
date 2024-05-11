package xmiddlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"strings"
	"time"
	"xcore/common/xtoken"
	"xcore/core/xvariable"
)

type Resp struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LogMiddleHandler  该函数前需要调用 tokenCheck 插件
func LogMiddleHandler(c *gin.Context) {
	respWriter := &ResponseWriterWrapper{c.Writer, bytes.NewBuffer([]byte{})}
	// 获取基本日志变量
	c.Writer = respWriter
	startTime := time.Now().UnixMilli()

	c.Next()
	cosTime := time.Now().UnixMilli() - startTime
	payload := xtoken.GetBindCustomPayload(c)
	// 响应状态码
	//responseStatus := int64(c.Writer.Status())

	respResult := (*respWriter.Body).String()

	var resp Resp
	_ = json.Unmarshal([]byte(respResult), &resp)

	//fmt.Printf("[%s]【%s】 %s CosTime:%dms RespStatus:%d Resp:%s\n", time.Now().Format(time.DateTime), c.Request.Method, c.Request.URL.String(), cosTime, responseStatus, resp.Data)
	xvariable.Logger.InfoLog.InfoContext(c, resp.Msg, slog.Int64("cosTime", cosTime), slog.String("path", c.FullPath()), slog.String("ip", c.RemoteIP()), slog.String("operateBy", payload.NickName), slog.String("method", c.Request.Method), slog.Int64("code", resp.Code), slog.String("data", strings.ReplaceAll(fmt.Sprintf("%s", resp.Data), "\\", "")))
}

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
