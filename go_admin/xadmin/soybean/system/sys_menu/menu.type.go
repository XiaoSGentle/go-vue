package sys_menu

import baseType "xadmin/soybean/dao/model/base"

type SysMenuListResp struct {
	baseType.PageResult
	Records []SysMenuList `json:"records"`
}

type AddOrUpDateSysMenuParam struct {
	ParentId   int32  `json:"parentId" form:"parentId" zh_comment:"上级目录" en_comment:"parentId" validate:"gte=0"`
	MenuType   string `json:"menuType" form:"menuType" zh_comment:"目录类型" en_comment:"menuType" validate:"required"`
	MenuName   string `json:"menuName" form:"menuName" zh_comment:"目录名称" en_comment:"menuName" validate:"required"`
	RouteName  string `json:"routeName" form:"routeName" zh_comment:"路由名称" en_comment:"routeName" validate:"required"`
	RoutePath  string `json:"routePath" form:"routePath" zh_comment:"路由地址" en_comment:"routePath" validate:"required"`
	Component  string `json:"component" form:"component" zh_comment:"组件" en_comment:"component" validate:"required"`
	HideInMenu bool   `json:"hideInMenu"`
	Order      int32  `json:"order"`
	I18NKey    string `json:"i18nKey"`
	Icon       string `json:"icon"`
	IconType   string `json:"iconType"`
	Status     string `json:"status"`
}

type SysMenuList struct {
	baseType.BaseRecord
	ParentId   int32         `json:"parentId"`
	MenuType   string        `json:"menuType"`
	MenuName   string        `json:"menuName"`
	RouteName  string        `json:"routeName"`
	RoutePath  string        `json:"routePath"`
	Component  string        `json:"component"`
	HideInMenu bool          `json:"hideInMenu"`
	Order      int32         `json:"order"`
	I18NKey    string        `json:"i18nKey"`
	Icon       string        `json:"icon"`
	IconType   string        `json:"iconType"`
	Children   []SysMenuList `json:"children"`
}
