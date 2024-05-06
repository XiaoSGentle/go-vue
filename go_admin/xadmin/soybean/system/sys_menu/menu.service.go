package sys_menu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"xadmin/soybean/dao/model"
	baseType "xadmin/soybean/dao/model/base"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
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
	metaMenuHide := "2"
	if param.HideInMenu {
		metaMenuHide = "1"
	}
	updates, err := menuQuery.WithContext(c).Where(menuQuery.ID.Eq(id)).Updates(&model.SysMenu{
		Name:       param.MenuName,
		RouterName: param.RouteName,
		Path:       param.RoutePath,
		ParentID:   param.ParentId,
		Component:  param.Component,
		//Props:            param.,
		Status:           param.Status,
		Type:             param.MenuType,
		MetaIconType:     param.IconType,
		MetaOrder:        param.Order,
		MetaConstant:     2,
		MetaHideInMenu:   metaMenuHide,
		MetaRequiresAuth: 1,
		MetaIcon:         param.Icon,
		MetaLocalIcon:    "",
		MetaI18nKey:      param.I18NKey,
		//MetaHref:         "",
		//MetaKeepAlive:    0,
		MetaTitle: param.RouteName,
		//MetaActiveMenu: "",

		//MetaFixedInTab: 0,
		MetaQuery:     "",
		Version:       0,
		SoftDeleteTag: 0,
		UpdateTime:    time.Now(),
		UpdateUID:     0,
		CreateUID:     0,
		CreateBy:      "",
		CreateTime:    time.Now(),
		UpdateBy:      "",
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
	err = sysMenuQuery.Create(&model.SysMenu{
		Name:       param.MenuName,
		RouterName: param.RouteName,
		Path:       param.RoutePath,
		ParentID:   param.ParentId,
		Component:  param.Component,
		//Props:            param.,
		Status:           param.Status,
		Type:             param.MenuType,
		MetaIconType:     param.IconType,
		MetaOrder:        param.Order,
		MetaConstant:     2,
		MetaHideInMenu:   "2",
		MetaRequiresAuth: 1,
		MetaIcon:         param.Icon,
		MetaLocalIcon:    "",
		MetaI18nKey:      param.I18NKey,
		//MetaHref:         "",
		//MetaKeepAlive:    0,
		MetaTitle: param.RouteName,
		//MetaActiveMenu: "",
		MetaMultiTab: "0",
		//MetaFixedInTab: 0,
		MetaQuery:     "",
		Version:       0,
		SoftDeleteTag: 0,
		UpdateTime:    time.Now(),
		UpdateUID:     0,
		CreateUID:     0,
		CreateBy:      "",
		CreateTime:    time.Now(),
		UpdateBy:      "",
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
		PageResult: baseType.PageResult{
			PageParam: baseType.PageParam{
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
			BaseRecord: baseType.BaseRecord{
				ID:         menu.ID,
				CreateBy:   menu.CreateBy,
				CreateTime: menu.CreateTime.String(),
				UpdateBy:   menu.UpdateBy,
				UpdateTime: menu.UpdateTime.String(),
				Status:     menu.Status,
			},
			ParentId:   menu.ParentID,
			MenuType:   menu.Type,
			MenuName:   menu.Name,
			RouteName:  menu.RouterName,
			HideInMenu: menu.MetaHideInMenu == "1",
			RoutePath:  menu.Path,
			Component:  menu.Component,
			Order:      menu.MetaOrder,
			I18NKey:    menu.MetaI18nKey,
			Icon:       menu.MetaIcon,
			IconType:   menu.MetaIconType,
			Children:   nil,
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
