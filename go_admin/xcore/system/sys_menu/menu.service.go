package sys_menu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xcore/dao/model"
	baseType "xcore/dao/model/base"
	"xcore/dao/query"
)

type ISysMenuService interface {
	GetMenuList(c *gin.Context) (sysMenuListResp SysMenuListResp, err error)
}

func NewSysMenuService(db *gorm.DB) ISysMenuService {
	return &SysMenuService{db: db}
}

type SysMenuService struct {
	db *gorm.DB
}

func (s SysMenuService) GetMenuList(c *gin.Context) (sysMenuListResp SysMenuListResp, err error) {
	menuQuery := query.Use(s.db).SysMenu
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
			ParentId:  menu.ParentID,
			MenuType:  menu.Type,
			MenuName:  menu.Name,
			RouteName: menu.RouterName,
			RoutePath: menu.Path,
			Component: menu.Component,
			Order:     menu.MetaOrder,
			I18NKey:   menu.MetaI18nKey,
			Icon:      menu.MetaIcon,
			IconType:  menu.MetaIconType,
			Children:  nil,
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
