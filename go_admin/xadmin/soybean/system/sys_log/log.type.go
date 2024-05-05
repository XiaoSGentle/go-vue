package sys_log

type LoginParam struct {
	UserName string `json:"userName" zh_comment:"用户名" en_comment:"" validate:"required"`
	Password string `json:"password" zh_comment:"密码" en_comment:"" validate:"required"`
}
type RefreshTokenParam struct {
	RefreshToken string `json:"refreshToken" zh_comment:"LogFileList" en_comment:"LogFileList" validate:"required"`
}
type CheckCaptchaParam struct {
	Dots []fontDots `json:"dots" validate:"required" comment:"点"`
	Key  string     `json:"key" validate:"required" comment:"验证码唯一标识"`
}
type fontDots struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Index int `json:"index"`
}
type GetSmsTextParam struct {
	CheckCaptchaParam
	PhoneNum string `json:"phoneNum" validate:"required" comment:"手机号码"`
}

type SmsCodeLoginParam struct {
	PhoneNum string `json:"phoneNum" validate:"required" comment:"手机号码"`
	Code     string `json:"code" validate:"required" comment:"验证码"`
}

type Oauth2MaxKeyAccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type Oauth2MaxKeyUserInfo struct {
	Birthday       string `json:"birthday"`
	Gender         int    `json:"gender"`
	DisplayName    string `json:"displayName"`
	DepartmentId   string `json:"departmentId"`
	Mobile         string `json:"mobile"`
	CreateDate     string `json:"createdate"`
	Title          string `json:"title"`
	UserId         string `json:"userId"`
	OnlineTicket   string `json:"online_ticket"`
	EmployeeNumber string `json:"employeeNumber"`
	RealName       string `json:"realname"`
	Institution    string `json:"institution"`
	RandomId       string `json:"randomId"`
	State          string `json:"state"`
	Department     string `json:"department"`
	User           string `json:"user"`
	Email          string `json:"email"`
	Username       string `json:"username"`
}

type Oauth2Param struct {
	Code string `json:"code"`
}
type Oauth2BindParam struct {
	BindKey  string `json:"bindKey,omitempty"`
	UserName string `json:"username" validate:"required" comment:"用户名" json:"userName,omitempty"`
	Password string `json:"password" validate:"required" comment:"密码" json:"password,omitempty"`
}

type Oauth2BindRedisSaveParam struct {
	PlatformName     string `json:"platformName,omitempty"`
	PlatformUserUUID string `json:"platformUserUUID" json:"platformUserUUID,omitempty"`
}

type LoginVo struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type UserInfoVo struct {
	UserId   string   `json:"userId"`
	UserName string   `json:"userName"`
	Roles    []string `json:"roles"`
	Apis     []string `json:"apis"`
}
type RedirectVo struct {
	Url string `json:"url"`
}

type CaptchaVo struct {
	Code        int    `json:"code"`
	ImageBase64 string `json:"image_base64"`
	ThumbBase64 string `json:"thumb_base64"`
	CaptchaKey  string `json:"captcha_key"`
}
