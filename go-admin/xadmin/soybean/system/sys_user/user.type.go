package sys_user

import (
	"xcore/common/xtype/xbase"
)

type SysUserListParam struct {
	xbase.PageParam
}

type SysUserList struct {
	xbase.BaseRecord
	UserName      string   `json:"userName"`
	UserGender    string   `json:"userGender"`
	NickName      string   `json:"nickName"`
	UserPhone     string   `json:"userPhone"`
	UserEmail     string   `json:"userEmail"`
	UserRoles     []string `json:"userRoles"`
	LastOnLine    string   `json:"lastOnLine"`
	LastCpWd      string   `json:"lastCpWd"`
	NeedChangePwd string   `json:"needChangePwd"`
}

type AddOrUpdateSysUserParam struct {
	UserName   string   `json:"userName"`
	UserGender string   `json:"userGender"`
	NickName   string   `json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	Status     string   `json:"status"`
	UserRoles  []string `json:"userRoles"`
}

type SysUserListResp struct {
	xbase.PageResult
	Records []SysUserList `json:"records"`
}
