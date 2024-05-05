package sys_dict

import baseType "xadmin/soybean/dao/model/base"

type SysDictListParam struct {
	baseType.PageParam
	Code string `json:"code"`
}

type SysDictListResp struct {
	baseType.PageResult
	Records []SysDictList `json:"records"`
}
type SysDictDataListParam struct {
	baseType.PageParam
	Code string `json:"code"`
}
type SysDictDataListResp struct {
	baseType.PageResult
	Records []SysDictDataList `json:"records"`
}

type AddOrUpDateSysDictDataParam struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Sort     int32  `json:"sort"`
	TypeCode string `json:"typeCode"`
	Status   string `json:"status"`
}

type AddOrUpDateSysDictParam struct {
	Name        string `json:"dictName"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SysDictList struct {
	baseType.BaseRecord
	Name        string `json:"dictName"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
type SysDictDataList struct {
	baseType.BaseRecord
	Label string `json:"label"`
	Value string `json:"value"`
	Sort  int32  `json:"sort"`
}

type DictInfo struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Sort  int32  `json:"sort"`
}
