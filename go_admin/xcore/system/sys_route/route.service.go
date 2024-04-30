package sys_route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"xcore/dao/model"
	"xcore/dao/query"
)

type IRouteService interface {
	GetUserRouters(c *gin.Context) (getUserRoutersVo GetUserRoutersVo, err error)
	GetConstantRoutes(c *gin.Context) (getUserRoutersVo []UserRouter, err error)
}

func NewRouteService(db *gorm.DB) IRouteService {
	return &RouteService{db: db}
}

type RouteService struct {
	db *gorm.DB
}

func (r RouteService) GetUserRouters(c *gin.Context) (getUserRoutersVo GetUserRoutersVo, err error) {
	menuQuery := query.Use(r.db).WithContext(c).SysMenu

	find, err := menuQuery.Find()
	if err != nil {
		return GetUserRoutersVo{}, err
	}
	result := GetUserRoutersVo{
		Home:   "/",
		Routes: SysMenuToRouterVoTree(find),
	}

	return result, nil
}
func (r RouteService) GetConstantRoutes(c *gin.Context) (getUserRoutersVo []UserRouter, err error) {
	return []UserRouter{
		{
			Name:      "login",
			Path:      "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			Component: "layout.blank$view.login",
			RouterMeta: UserRouterMeta{
				Constant:   true,
				HideInMenu: true,
				I18nKey:    "route.login",
				Title:      "login",
			},
		},
	}, nil
}

func SysMenuToRouterVoTree(sysMenuList []*model.SysMenu) (menuVoList []UserRouter) {
	for _, menu := range sysMenuList {
		m := UserRouter{
			ID:        fmt.Sprintf("%d", menu.ID),
			PID:       menu.ParentID,
			Name:      menu.Name,
			Path:      menu.Path,
			Component: menu.Component,
			RouterMeta: UserRouterMeta{
				IconType:     menu.MetaIconType,
				Order:        menu.MetaOrder,
				Constant:     menu.MetaConstant == 1,
				HideInMenu:   menu.MetaHideInMenu == 1,
				RequiresAuth: menu.MetaRequiresAuth == 1,
				Icon:         menu.MetaIcon,
				LocalIcon:    menu.MetaLocalIcon,
				I18nKey:      menu.MetaI18nKey,
				Href:         menu.MetaHref,
				KeepAlive:    menu.MetaKeepAlive == 1,
				Title:        menu.MetaTitle,
				ActiveMenu:   menu.MetaActiveMenu,
				MultiTab:     menu.MetaMultiTab == 1,
				FixedInTab:   menu.MetaFixedInTab,
				Query:        "",
			},
			Children: nil,
		}
		menuVoList = append(menuVoList, m)
	}
	menuVoList = listToTree(menuVoList, 0)
	return
}

func listToTree(list []UserRouter, Pid int32) (tree []UserRouter) {
	res := make([]UserRouter, 0)
	for _, v := range list {
		if v.PID == Pid {
			num, _ := strconv.Atoi(v.ID)
			v.Children = listToTree(list, int32(num))
			res = append(res, v)
		}
	}
	tree = res
	return
}
