package sys_role

import (
	baseType "xcore/dao/model/base"
)

type SysRoleListParam struct {
	baseType.PageParam
}

type SysRoleList struct {
	baseType.BaseRecord
	/** role name */
	RoleName string `json:"roleName"`
	/** role code */
	RoleCode string `json:"roleCode"`
	/** role description */
	RoleDesc string `json:"roleDesc"`
}

type SysRoleListResp struct {
	baseType.PageResult
	Records []SysRoleList `json:"records"`
}
