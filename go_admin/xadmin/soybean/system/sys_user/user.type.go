package sys_user

import (
	baseType "xadmin/soybean/dao/model/base"
)

type SysUserListParam struct {
	baseType.PageParam
}

type SysUserList struct {
	baseType.BaseRecord
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
	baseType.PageResult
	Records []SysUserList `json:"records"`
}
