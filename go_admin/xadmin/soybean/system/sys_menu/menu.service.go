package sys_menu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
	"xcore/common/xtoken"
	"xcore/common/xtype/xbase"
	"xcore/common/xtype/xbool"
)

type ISysMenuService interface {
	GetMenuList(c *gin.Context) (resp SysMenuListResp, err error)
	AddMenu(c *gin.Context, param *AddOrUpDateSysMenuParam) (err error)
	UpdateMenu(c *gin.Context, id int32, param *AddOrUpDateSysMenuParam) (err error)
	DeleteMenu(c *gin.Context, ids []int32) (err error)
}

func NewSysMenuService(db *gorm.DB) ISysMenuService {
	return &SysMenuService{db: db, query: query.Use(db)}
}

type SysMenuService struct {
	db    *gorm.DB
	query *query.Query
}

func (m SysMenuService) DeleteMenu(c *gin.Context, ids []int32) (err error) {
	menuQuery := m.query.SysMenu
	info, err := menuQuery.WithContext(c).Where(menuQuery.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.Error != nil {
		return info.Error
	}
	if info.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}

func (m SysMenuService) UpdateMenu(c *gin.Context, id int32, param *AddOrUpDateSysMenuParam) (err error) {
	menuQuery := m.query.SysMenu
	payload := xtoken.GetBindCustomPayload(c)
	updates, err := menuQuery.WithContext(c).Where(menuQuery.ID.Eq(id)).Updates(&model.SysMenu{
		Name:           param.MenuName,
		RouterName:     param.RouteName,
		Path:           param.RoutePath,
		ParentID:       param.ParentId,
		Component:      param.Component,
		Status:         param.Status,
		Type:           param.MenuType,
		MetaIconType:   param.IconType,
		MetaOrder:      param.Order,
		MetaConstant:   xbool.BooleanTo(param.Constant, "1", "2"),
		MetaHideInMenu: xbool.BooleanTo(param.HideInMenu, "1", "2"),
		MetaIcon:       param.Icon,
		MetaLocalIcon:  "",
		MetaI18nKey:    param.I18NKey,
		MetaHref:       param.Href,
		MetaKeepAlive:  xbool.BooleanTo(param.KeepAlive, "1", "2"),
		MetaTitle:      param.RouteName,
		MetaActiveMenu: param.ActiveMenu,
		MetaMultiTab:   xbool.BooleanTo(param.MultiTab, "1", "2"),
		MetaFixedInTab: param.FixedIndexInTab,
		MetaQuery:      strings.Join(param.Query, ","),
		UpdateTime:     time.Now(),
		UpdateUID:      payload.Uid,
		UpdateBy:       payload.NickName,
	})
	if err != nil {
		return err
	}
	if updates.Error != nil {
		return updates.Error
	}
	return
}

func (m SysMenuService) AddMenu(c *gin.Context, param *AddOrUpDateSysMenuParam) (err error) {
	sysMenuQuery := m.query.SysMenu.WithContext(c)
	payload := xtoken.GetBindCustomPayload(c)
	err = sysMenuQuery.Create(&model.SysMenu{
		Name:           param.MenuName,
		RouterName:     param.RouteName,
		Path:           param.RoutePath,
		ParentID:       param.ParentId,
		Component:      param.Component,
		Status:         param.Status,
		Type:           param.MenuType,
		MetaIconType:   param.IconType,
		MetaOrder:      param.Order,
		MetaConstant:   xbool.BooleanTo(param.Constant, "1", "2"),
		MetaHideInMenu: xbool.BooleanTo(param.HideInMenu, "1", "2"),
		MetaIcon:       xbool.BooleanTo(param.MenuType == "1", param.Icon, ""),
		MetaLocalIcon:  xbool.BooleanTo(param.MenuType == "1", "", param.Icon),
		MetaI18nKey:    param.I18NKey,
		MetaHref:       param.Href,
		MetaKeepAlive:  xbool.BooleanTo(param.KeepAlive, "1", "2"),
		MetaTitle:      param.RouteName,
		MetaActiveMenu: param.ActiveMenu,
		MetaMultiTab:   xbool.BooleanTo(param.MultiTab, "1", "2"),
		MetaFixedInTab: param.FixedIndexInTab,
		MetaQuery:      strings.Join(param.Query, ","),
		CreateTime:     time.Now(),
		CreateUID:      payload.Uid,
		CreateBy:       payload.NickName,
	})
	return
}

func (m SysMenuService) GetMenuList(c *gin.Context) (resp SysMenuListResp, err error) {
	menuQuery := query.Use(m.db).SysMenu
	find, err := menuQuery.Find()
	if err != nil {
		return SysMenuListResp{}, err
	}

	result := sysMenuToSysMenuListRespTree(find)

	return SysMenuListResp{
		PageResult: xbase.PageResult{
			PageParam: xbase.PageParam{
				Current: 1,
				Size:    10,
			},
			Total: int64(len(result)),
		},
		Records: result,
	}, nil
}

func sysMenuToSysMenuListRespTree(sysMenuList []*model.SysMenu) (menuVoList []SysMenuList) {
	for _, menu := range sysMenuList {
		m := SysMenuList{
			BaseRecord: xbase.BaseRecord{
				ID:         menu.ID,
				CreateBy:   menu.CreateBy,
				CreateTime: menu.CreateTime.String(),
				UpdateBy:   menu.UpdateBy,
				UpdateTime: menu.UpdateTime.String(),
				Status:     menu.Status,
			},
			ParentId:        menu.ParentID,
			MenuType:        menu.Type,
			MenuName:        menu.Name,
			RouteName:       menu.RouterName,
			RoutePath:       menu.Path,
			Component:       menu.Component,
			HideInMenu:      menu.MetaHideInMenu == "1",
			Order:           menu.MetaOrder,
			I18NKey:         menu.MetaI18nKey,
			Icon:            menu.MetaIcon,
			IconType:        menu.MetaIconType,
			Status:          menu.Status,
			KeepAlive:       menu.MetaKeepAlive == "1",
			Constant:        menu.MetaConstant == "1",
			ActiveMenu:      menu.MetaActiveMenu,
			FixedIndexInTab: menu.MetaFixedInTab,
			MultiTab:        menu.MetaMultiTab == "1",
			Href:            menu.MetaHref,
			Query:           strings.Split(menu.MetaQuery, ","),
			Children:        []SysMenuList{},
		}
		menuVoList = append(menuVoList, m)
	}
	menuVoList = userRouterListToTree(menuVoList, 0)
	return
}
func userRouterListToTree(list []SysMenuList, Pid int32) (tree []SysMenuList) {
	res := make([]SysMenuList, 0)
	for _, v := range list {
		if (v.ParentId) == (Pid) {
			v.Children = userRouterListToTree(list, v.ID)
			res = append(res, v)
		}
	}
	tree = res
	return
}
