package sys_route

type GetUserRoutersVo struct {
	Routes []UserRouter `json:"routes,omitempty"`
	Home   string       `json:"home,omitempty"`
}

type UserRouter struct {
	ID         string         `json:"id,omitempty"`
	PID        int32          `json:"-"`
	Name       string         `json:"name,omitempty"`
	Path       string         `json:"path,omitempty"`
	Props      bool           `json:"props,omitempty"`
	Component  string         `json:"component,omitempty"`
	RouterMeta UserRouterMeta `json:"meta,omitempty"`
	Children   []UserRouter   `json:"children,omitempty"`
}

type UserRouterMeta struct {
	Title      string           `json:"title,omitempty"`
	I18nKey    string           `json:"i18nKey,omitempty"`
	Order      int32            `json:"order,omitempty"`
	Role       []string         `json:"role,omitempty"`
	Constant   bool             `json:"constant,omitempty"`
	KeepAlive  bool             `json:"keepAlive,omitempty"`
	Icon       string           `json:"icon,omitempty"`
	LocalIcon  string           `json:"localIcon,omitempty"`
	Href       string           `json:"href,omitempty"`
	HideInMenu bool             `json:"hideInMenu,omitempty"`
	ActiveMenu string           `json:"activeMenu,omitempty"`
	MultiTab   bool             `json:"multiTab,omitempty"`
	FixedInTab int32            `json:"fixedInTab,omitempty"`
	Query      []MenuQueryParam `json:"query,omitempty"`
}
type MenuQueryParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type RouterTreeSimpleResp struct {
	ID       string                 `json:"id"`
	Label    string                 `json:"label"`
	PID      string                 `json:"pId"`
	Children []RouterTreeSimpleResp `json:"children,omitempty"`
}

type AllApisResp struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AllRolesResp struct {
	Id       string `json:"id"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
}
