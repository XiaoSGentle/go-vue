package xauth

import (
	"fmt"
	"gorm.io/gorm"
	"xcore/common/xcache"
	"xcore/common/xtype/xslice"
)

var (
	AuthApisPermitKey  = "AUTH_APIS_PERMIT_KEY::"
	AuthMenusPermitKey = "AUTH_MENUS_PERMIT_KEY::"
)

type Auth struct {
	db             *gorm.DB
	apiCacheStore  *xcache.CacheStore[[]string]
	menuCacheStore *xcache.CacheStore[[]string]
}
type AuthContent struct {
	RoleCode string
	Menus    []string
	Apis     []string
}

func NewAuth(permits []AuthContent) *Auth {
	var auth = new(Auth)
	auth.apiCacheStore = xcache.NewCacheStore[[]string](0, AuthApisPermitKey)
	auth.menuCacheStore = xcache.NewCacheStore[[]string](0, AuthMenusPermitKey)
	for _, permit := range permits {
		auth.menuCacheStore.Set(permit.RoleCode, permit.Menus)
		auth.apiCacheStore.Set(permit.RoleCode, permit.Apis)
	}

	return auth
}

func (r *Auth) CheckApiPermit(roles []string, requestMethod, apiPath string) bool {
	var allApis []string
	for _, role := range roles {
		allApis = append(allApis, r.apiCacheStore.Get(role)...)
	}
	return xslice.StringExist(allApis, fmt.Sprintf("%s::%s", requestMethod, apiPath))
}

func (r *Auth) GetMenuIdsByRoles(roles []string) []string {
	var allMenusIds []string
	for _, role := range roles {
		allMenusIds = append(allMenusIds, r.menuCacheStore.Get(role)...)
	}
	return xslice.StringDuplicate(allMenusIds)
}

func (r *Auth) GetApiCodesByRoles(roles []string) []string {
	var allApiCodes []string
	for _, role := range roles {
		allApiCodes = append(allApiCodes, r.apiCacheStore.Get(role)...)
	}
	return xslice.StringDuplicate(allApiCodes)
}

func (r *Auth) RefreshAuthStore(permits []AuthContent) {
	_ = r.apiCacheStore.Clear()
	_ = r.menuCacheStore.Clear()
	for _, permit := range permits {
		r.menuCacheStore.Set(permit.RoleCode, permit.Menus)
		r.apiCacheStore.Set(permit.RoleCode, permit.Apis)
	}
	return
}
