package sys_route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"xadmin/soybean/dao/model"
	"xcore/common/xtype/xbool"
	"xcore/core/xconst"
)

type IRouteService interface {
	IsRouteExist(c *gin.Context, routerName string) (exist bool, err error)
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

func (r RouteService) IsRouteExist(c *gin.Context, routerName string) (exist bool, err error) {
	menuQuery := r.db.WithContext(c).Model(model.SysMenu{})
	var routerCount int64
	err = menuQuery.Where("name = ?", routerName).Count(&routerCount).Error
	exist = routerCount > 0
	return
}

func (r RouteService) GetAllRoles(c *gin.Context) (allRoles []AllRolesResp, err error) {
	var find []model.SysRole
	db := r.db.WithContext(c).Model(model.SysRole{}).Find(&find)
	err = db.Error
	for _, role := range find {
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
	var find []model.SysAPI
	db := r.db.WithContext(c).Model(model.SysAPI{}).Find(&find)
	err = db.Error
	for _, api := range find {
		result = append(result, AllApisResp{
			Code: api.APICode,
			Name: api.APICode,
		})
	}
	return result, nil
}

func (r RouteService) GetMenuTreeSimple(c *gin.Context) (routerTreeSimpleResp []RouterTreeSimpleResp, err error) {
	var find []model.SysMenu
	db := r.db.WithContext(c).Model(model.SysMenu{}).Find(&find)
	err = db.Error
	result := sysMenuToSimpleRouterTree(find)
	return result, nil
}

func (r RouteService) GetAllPages(c *gin.Context, roles []string) (allPages []string, err error) {
	var find []model.SysMenu
	db := r.db.WithContext(c).Model(model.SysMenu{}).Find(&find)
	err = db.Error
	for _, menu := range find {
		allPages = append(allPages, menu.RouterName)
	}
	return
}

func (r RouteService) GetUserRouters(c *gin.Context, ids []int32) (getUserRoutersVo GetUserRoutersVo, err error) {
	var menuInSql []model.SysMenu
	err = r.db.WithContext(c).Model(model.SysMenu{}).
		Where("id in ?", ids).
		Where("status = ?", xconst.StatusOK).
		Where("delete_tag = ?", xconst.NotDelete).
		Find(&menuInSql).Error
	if err != nil {
		return GetUserRoutersVo{}, err
	}
	result := GetUserRoutersVo{
		Home:   "home",
		Routes: sysMenuToRouterVoTree(menuInSql),
	}
	return result, nil
}
func (r RouteService) GetConstantRoutes(c *gin.Context) (getUserRoutersVo []UserRouter, err error) {
	return []UserRouter{
		{
			Name:      "login",
			Props:     true,
			Path:      "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			Component: "layout.blank$view.login",
			RouterMeta: UserRouterMeta{
				HideInMenu: true,
				Constant:   true,
				I18nKey:    "route.login",
				Title:      "login",
			},
		},
	}, nil
}

func sysMenuToRouterVoTree(sysMenuList []model.SysMenu) (menuVoList []UserRouter) {
	for _, menu := range sysMenuList {
		m := UserRouter{
			ID:        fmt.Sprintf("%d", menu.ID),
			PID:       menu.ParentID,
			Name:      menu.RouterName,
			Path:      menu.Path,
			Component: menu.Component,
			RouterMeta: UserRouterMeta{
				Order:      menu.MetaOrder,
				HideInMenu: menu.MetaHideInMenu == xconst.StatusOK,
				Icon:       xbool.BooleanTo(menu.MetaIconType == xconst.StatusOK, menu.MetaIcon, ""),
				LocalIcon:  xbool.BooleanTo(menu.MetaIconType == xconst.StatusBanned, menu.MetaIcon, ""),
				I18nKey:    menu.MetaI18nKey,
				Href:       menu.MetaHref,
				KeepAlive:  menu.MetaKeepAlive == xconst.StatusOK,
				Title:      menu.MetaTitle,
				ActiveMenu: menu.MetaActiveMenu,
				MultiTab:   menu.MetaMultiTab == xconst.StatusOK,
				FixedInTab: menu.MetaFixedInTab,
				Query:      []MenuQueryParam{},
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

func sysMenuToSimpleRouterTree(sysMenuList []model.SysMenu) (routerTreeSimpleResp []RouterTreeSimpleResp) {
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
