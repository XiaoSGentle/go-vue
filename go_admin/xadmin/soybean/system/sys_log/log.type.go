package sys_log

import baseType "xadmin/soybean/dao/model/base"

type SysLogListResp struct {
	baseType.PageResult
	Records []string `json:"records"`
}
