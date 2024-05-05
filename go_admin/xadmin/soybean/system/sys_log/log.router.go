package sys_log

import (
	"github.com/gin-gonic/gin"
	"os"
	"xcore/common/xmiddlewares"
	"xcore/core/xcore"
)

var LogGroup = xcore.Group("/system/log", newSysLogHandler, regLog, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)

func regLog(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *logHandler) {
		rg.GET("/list", handle.LogFileList)
		rg.GET("/:fileName", handle.DownLogFile)
	})
}
func newSysLogHandler() *logHandler {
	return &logHandler{}
}

type logHandler struct {
}

func (h logHandler) LogFileList(c *gin.Context) {
	list := GetFiles("./logs")
	for _, s := range list {
		info, err := s.Info()
		if err != nil {
			return
		}
		// todo:写成接口
		println(info.Size())
		println(s.Name())
	}
}

func (h logHandler) DownLogFile(c *gin.Context) {

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
