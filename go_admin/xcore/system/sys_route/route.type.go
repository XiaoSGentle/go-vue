package sys_route

type GetUserRoutersVo struct {
	Routes []UserRouter `json:"routes,omitempty"`
	Home   string       `json:"home,omitempty"`
}

type UserRouter struct {
	ID         string         `json:"id"`
	PID        int32          `json:"-"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Component  string         `json:"component"`
	RouterMeta UserRouterMeta `json:"meta"`
	Children   []UserRouter   `json:"children,omitempty"`
}

type UserRouterMeta struct {
	IconType     string `json:"iconType"`
	Order        int32  `json:"order"`
	Constant     bool   `json:"constant"`
	HideInMenu   bool   `json:"hideInMenu"`
	RequiresAuth bool   `json:"requiresAuth"`
	Icon         string `json:"icon"`
	LocalIcon    string `json:"localIcon"`
	I18nKey      string `json:"i18nKey"`
	Href         string `json:"href,omitempty"`
	KeepAlive    bool   `json:"keepAlive"`
	Title        string `json:"title"`
	ActiveMenu   string `json:"activeMenu"`
	MultiTab     bool   `json:"multiTab"`
	FixedInTab   int32  `json:"fixedInTab"`
	Query        string `json:"query"`
}
