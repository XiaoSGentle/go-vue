package sys_log

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"xcore/common/xerror"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	baseType "xcore/common/xtype/xbase"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
)

var LogGroup = xcore.Group("/system/log", newSysLogHandler, regLog, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)
var NoLogLogGroup = xcore.Group("/system/log", newSysLogHandler, noLogLogGroup, xmiddlewares.Authorize)

func regLog(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *logHandler) {
		rg.GET("/list", handle.LogFileList)
		rg.GET("/download/:fileName", handle.DownLogFile)
	})
}

func noLogLogGroup(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *logHandler) {
		rg.GET("/:level", handle.LogContent)
	})
}
func newSysLogHandler() *logHandler {
	return &logHandler{}
}

type logHandler struct {
}

func (h logHandler) LogFileList(c *gin.Context) {
	list := GetFiles("./logs")
	var fileZips []SysLogFileZipsResp
	for _, s := range list {
		if strings.HasSuffix(s.Name(), ".log.gz") {
			stat, _ := os.Stat("./logs/" + s.Name())
			modTimer := stat.ModTime()
			info, _ := s.Info()
			fileSizeMB := float64(info.Size()) / (1024 * 1024)
			fileSizeMB = float64(int(fileSizeMB*100)) / 100
			fileZips = append(fileZips, SysLogFileZipsResp{
				FileName:   s.Name(),
				FileSize:   fmt.Sprintf("%.2fMB", fileSizeMB),
				CreateData: modTimer,
			})
		}
	}
	xresponse.SuccessCtx(c, fileZips)
}

func (h logHandler) DownLogFile(c *gin.Context) {
	var param struct {
		FileName string `json:"fileName" uri:"fileName" form:"fileName" zh_comment:"文件名称" en_comment:"file name" validate:"required"`
	}

	err := c.BindUri(&param)
	if err != nil {
		xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.PARAM_BIND_ERROR))
		return
	}
	err = xvalidate.ValidateStruct(&param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	// 获取文件路径
	filePath := fmt.Sprintf("./logs/%s", param.FileName)
	// 设置文件名和文件类型
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	// 返回文件
	c.File(filePath)
}

func (h logHandler) LogContent(c *gin.Context) {
	var level struct {
		Level string `json:"level" uri:"level" form:"level" zh_comment:"日志等级" en_comment:"log level" validate:"oneof=info error"`
	}
	err := c.BindUri(&level)

	if err != nil {
		xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.PARAM_BIND_ERROR))
		return
	}
	var pageInfo baseType.PageParam
	err = c.ShouldBind(&pageInfo)
	if err != nil {
		xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.PARAM_BIND_ERROR))
		return
	}
	err = xvalidate.ValidateStruct(&level)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	file, err := os.Open("./logs/" + level.Level + ".log")
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewReader(file)
	lines := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	var result SysLogListResp
	for i := len(lines) - (pageInfo.Size)*(pageInfo.Current-1) - 1; i >= len(lines)-pageInfo.Size*pageInfo.Current-1 && i >= 0; i-- {
		result.Records = append(result.Records, lines[i])
	}
	result.Size = pageInfo.Size
	result.Current = pageInfo.Current
	result.Total = int64(len(lines)/pageInfo.Size) + 1
	xresponse.SuccessCtx(c, result)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
func GetFiles(folder string) (filesList []os.DirEntry) {
	files, _ := os.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + "/" + file.Name())
		} else {
			filesList = append(filesList, file)
		}
	}
	return
}
