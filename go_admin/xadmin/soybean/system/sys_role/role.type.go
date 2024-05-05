package sys_role

import (
	baseType "xadmin/soybean/dao/model/base"
)

type SysRoleListParam struct {
	baseType.PageParam
}

type UpdateRoleHomeParam struct {
	RoleCode string `json:"roleCode"`
	Home     string `json:"home"`
}

type UpdateRolePermitParam struct {
	SysRolePermitResp
	RoleCode string `json:"roleCode"`
}

type SysRolePermitResp struct {
	ApiCodes []string `json:"apiCodes"`
	MenuIds  []string `json:"menuIds"`
}

type SysRoleList struct {
	baseType.BaseRecord
	/** role name */
	RoleName string `json:"roleName"`
	/** role code */
	RoleCode string `json:"roleCode"`
	/** role description */
	RoleDesc string `json:"roleDesc"`

	RoleHome string `json:"roleHome"`
}

type AddOrUpdateSysRoleParam struct {
	Status string `json:"status"`
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
