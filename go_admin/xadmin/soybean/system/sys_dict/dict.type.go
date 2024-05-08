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
	Code string `json:"code" form:"code" zh_comment:"当前页数" en_comment:"current" validate:"required"`
}
type SysDictDataListResp struct {
	baseType.PageResult
	Records []SysDictDataList `json:"records"`
}

type AddOrUpDateSysDictDataParam struct {
	Label  string `json:"label" zh_comment:"键" en_comment:"label" validate:"required"`
	Value  string `json:"value" zh_comment:"值" en_comment:"value" validate:"required"`
	Sort   int32  `json:"sort"`
	Code   string `json:"code" zh_comment:"字典类型" en_comment:"type code" validate:"required"`
	Status string `json:"status"`
}

type AddOrUpDateSysDictParam struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SysDictList struct {
	baseType.BaseRecord
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
type SysDictDataList struct {
	baseType.BaseRecord
	Label   string `json:"label"`
	Value   string `json:"value"`
	Code    string `json:"code"`
	EnLabel string `json:"enLabel"`
	Sort    int32  `json:"sort"`
}

type DictInfo struct {
	Label   string `json:"label"`
	EnLabel string `json:"enLabel"`
	Value   string `json:"value"`
	Sort    int32  `json:"sort"`
}
