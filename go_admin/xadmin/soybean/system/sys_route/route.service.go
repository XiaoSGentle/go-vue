package sys_route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
)

type IRouteService interface {
	GetUserRouters(c *gin.Context, ids []int32) (getUserRoutersVo GetUserRoutersVo, err error)
	GetConstantRoutes(c *gin.Context) (getUserRoutersVo []UserRouter, err error)
	GetAllPages(c *gin.Context, roles []string) (allPages []string, err error)
	GetMenuTreeSimple(c *gin.Context) (routerTreeSimpleResp []RouterTreeSimpleResp, err error)
	GetALLApis(c *gin.Context) (allApisResp []AllApisResp, err error)
	GetAllRoles(c *gin.Context) (allRoles []AllRolesResp, err error)
}

func NewRouteService(db *gorm.DB) IRouteService {
	return &RouteService{db: db}
}

type RouteService struct {
	db *gorm.DB
}

func (r RouteService) GetAllRoles(c *gin.Context) (allRoles []AllRolesResp, err error) {
	roleQuery := query.Use(r.db).SysRole
	findInSql, err := roleQuery.WithContext(c).Find()
	if err != nil {
		return []AllRolesResp{}, err
	}
	for _, role := range findInSql {
		allRoles = append(allRoles, AllRolesResp{
			Id:       fmt.Sprintf("%d", role.ID),
			RoleName: role.Name,
			RoleCode: role.Code,
		})
	}
	return
}

func (r RouteService) GetALLApis(c *gin.Context) (allApisResp []AllApisResp, err error) {
	var result []AllApisResp
	apiQuery := query.Use(r.db).SysAPI
	findApisInSql, err := apiQuery.WithContext(c).Find()
	if err != nil {
		return []AllApisResp{}, err
	}
	for _, api := range findApisInSql {
		result = append(result, AllApisResp{
			Code: api.APICode,
			Name: api.APICode,
		})
	}
	return result, nil
}

func (r RouteService) GetMenuTreeSimple(c *gin.Context) (routerTreeSimpleResp []RouterTreeSimpleResp, err error) {
	menuQuery := query.Use(r.db).WithContext(c).SysMenu
	find, err := menuQuery.Find()
	if err != nil {
		return []RouterTreeSimpleResp{}, err
	}
	result := sysMenuToSimpleRouterTree(find)
	return result, nil
}

func (r RouteService) GetAllPages(c *gin.Context, roles []string) (allPages []string, err error) {
	menuQuery := query.Use(r.db).WithContext(c).SysMenu
	menusInSql, err := menuQuery.Find()
	if err != nil {
		return []string{}, err
	}
	for _, menu := range menusInSql {
		allPages = append(allPages, menu.RouterName)
	}
	return
}

func (r RouteService) GetUserRouters(c *gin.Context, ids []int32) (getUserRoutersVo GetUserRoutersVo, err error) {
	menuQuery := query.Use(r.db).SysMenu
	find, err := menuQuery.WithContext(c).Where(menuQuery.ID.In(ids...)).Find()
	if err != nil {
		return GetUserRoutersVo{}, err
	}
	result := GetUserRoutersVo{
		Home:   "home",
		Routes: sysMenuToRouterVoTree(find),
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

func sysMenuToRouterVoTree(sysMenuList []*model.SysMenu) (menuVoList []UserRouter) {

	for _, menu := range sysMenuList {
		m := UserRouter{
			ID:        fmt.Sprintf("%d", menu.ID),
			PID:       menu.ParentID,
			Name:      menu.RouterName,
			Path:      menu.Path,
			Component: menu.Component,
			RouterMeta: UserRouterMeta{
				IconType:     menu.MetaIconType,
				Order:        menu.MetaOrder,
				Constant:     menu.MetaConstant == 1,
				HideInMenu:   menu.MetaHideInMenu == "1",
				RequiresAuth: menu.MetaRequiresAuth == 1,
				Icon:         menu.MetaIcon,
				LocalIcon:    menu.MetaLocalIcon,
				I18nKey:      menu.MetaI18nKey,
				Href:         menu.MetaHref,
				KeepAlive:    menu.MetaKeepAlive == "1",
				Title:        menu.MetaTitle,
				ActiveMenu:   menu.MetaActiveMenu,
				MultiTab:     menu.MetaMultiTab == "1",
				FixedInTab:   menu.MetaFixedInTab,
				Query:        "",
			},
			Children: nil,
		}
		menuVoList = append(menuVoList, m)
	}
	menuVoList = userRouterListToTree(menuVoList, 0)
	return
}
func userRouterListToTree(list []UserRouter, Pid int32) (tree []UserRouter) {

	res := make([]UserRouter, 0)
	for _, v := range list {
		if v.PID == Pid {
			num, _ := strconv.Atoi(v.ID)
			v.Children = userRouterListToTree(list, int32(num))
			res = append(res, v)
		}
	}
	tree = res
	return
}

func sysMenuToSimpleRouterTree(sysMenuList []*model.SysMenu) (routerTreeSimpleResp []RouterTreeSimpleResp) {
	for _, menu := range sysMenuList {
		routerTreeSimpleResp = append(routerTreeSimpleResp, RouterTreeSimpleResp{
			ID:       fmt.Sprintf("%d", menu.ID),
			Label:    menu.MetaI18nKey,
			PID:      fmt.Sprintf("%d", menu.ParentID),
			Children: nil,
		})
	}
	routerTreeSimpleResp = routerSimpleToTree(routerTreeSimpleResp, "0")
	return
}
func routerSimpleToTree(list []RouterTreeSimpleResp, Pid string) (tree []RouterTreeSimpleResp) {
	res := make([]RouterTreeSimpleResp, 0)
	for _, v := range list {
		if v.PID == Pid {
			num, _ := strconv.Atoi(v.ID)
			v.Children = routerSimpleToTree(list, strconv.Itoa(num))
			res = append(res, v)
		}
	}
	tree = res
	return
}
