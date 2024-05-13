package sys_log

import (
	"time"
	baseType "xcore/common/xtype/xbase"
)

type SysLogListResp struct {
	baseType.PageResult
	Records []string `json:"records"`
}

type SysLogFileZipsResp struct {
	FileName   string    `json:"fileName"`
	FileSize   string    `json:"fileSize"`
	CreateData time.Time `json:"createData"`
}
