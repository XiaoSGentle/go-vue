package sys_menu

import (
	baseType "xcore/common/xtype/xbase"
)

type SysMenuListResp struct {
	baseType.PageResult
	Records []SysMenuList `json:"records"`
}

type AddOrUpDateSysMenuParam struct {
	ParentId        int32           `json:"parentId" form:"parentId" zh_comment:"上级目录" en_comment:"parentId" validate:"gte=0"`
	MenuType        string          `json:"menuType" form:"menuType" zh_comment:"目录类型" en_comment:"menuType" validate:"oneof='1' '2'"`
	MenuName        string          `json:"menuName" form:"menuName" zh_comment:"目录名称" en_comment:"menuName" validate:"required"`
	RouteName       string          `json:"routeName" form:"routeName" zh_comment:"路由名称" en_comment:"routeName" validate:"required"`
	RoutePath       string          `json:"routePath" form:"routePath" zh_comment:"路由地址" en_comment:"routePath" validate:"required"`
	Component       string          `json:"component" form:"component" zh_comment:"组件" en_comment:"component"`
	HideInMenu      bool            `json:"hideInMenu"`
	Order           int32           `json:"order"`
	I18NKey         string          `json:"i18nKey"`
	Icon            string          `json:"icon"`
	IconType        string          `json:"iconType"`
	Status          string          `json:"status" zh_comment:"状态" en_comment:"status" validate:"oneof='1' '2'"`
	KeepAlive       bool            `json:"keepAlive"`
	Constant        bool            `json:"constant"`
	ActiveMenu      string          `json:"activeMenu"`
	FixedIndexInTab int32           `json:"fixedIndexInTab"`
	MultiTab        bool            `json:"multiTab"`
	Href            string          `json:"href"`
	PathParam       string          `json:"pathParam"`
	Query           []MenuQueryType `json:"query"`
}

type MenuQueryType struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SysMenuList struct {
	baseType.BaseRecord
	ParentId        int32         `json:"parentId"`
	MenuType        string        `json:"menuType"`
	MenuName        string        `json:"menuName"`
	RouteName       string        `json:"routeName"`
	RoutePath       string        `json:"routePath"`
	Component       string        `json:"component"`
	HideInMenu      bool          `json:"hideInMenu"`
	Order           int32         `json:"order"`
	I18NKey         string        `json:"i18nKey"`
	Icon            string        `json:"icon"`
	IconType        string        `json:"iconType"`
	Status          string        `json:"status"`
	KeepAlive       bool          `json:"keepAlive"`
	Constant        bool          `json:"constant"`
	ActiveMenu      string        `json:"activeMenu"`
	FixedIndexInTab int32         `json:"fixedIndexInTab"`
	MultiTab        bool          `json:"multiTab"`
	Href            string        `json:"href"`
	PathParam       string        `json:"pathParam"`
	Query           []string      `json:"query"`
	Children        []SysMenuList `json:"children"`
}
