package sys_menu

import baseType "xcore/dao/model/base"

type SysMenuListResp struct {
	baseType.PageResult
	Records []SysMenuList `json:"records"`
}

type SysMenuList struct {
	baseType.BaseRecord
	ParentId  int32         `json:"parentId"`
	MenuType  string        `json:"menuType"`
	MenuName  string        `json:"menuName"`
	RouteName string        `json:"routeName"`
	RoutePath string        `json:"routePath"`
	Component string        `json:"component"`
	Order     int32         `json:"order"`
	I18NKey   string        `json:"i18nKey"`
	Icon      string        `json:"icon"`
	IconType  string        `json:"iconType"`
	Children  []SysMenuList `json:"children"`
}
