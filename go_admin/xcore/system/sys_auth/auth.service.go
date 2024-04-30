package sys_auth

import (
	"gorm.io/gorm"
)

type IAuthService interface {
	UserLogin(param *LoginParam) (loginVo *LoginVo, err error)
	SmsLogin(phoneNum *string) (loginVo *LoginVo, err error)
	Oauth2Login(otherPlantFormUserKey string, plantFormKey string, effectiveTime int64) (loginVo *LoginVo, statesCode string, bindKey string, err error)
}

func NewAuthService(db *gorm.DB) IAuthService {
	return &AuthService{db: db}
}

type AuthService struct {
	db *gorm.DB
}

func (a AuthService) SmsLogin(phoneNum *string) (loginVo *LoginVo, err error) {
	return
}

func (a AuthService) UserLogin(param *LoginParam) (loginVo *LoginVo, err error) {
	return
}

func (a AuthService) Oauth2Login(otherPlantFormUserKey string, plantFormKey string, effectiveTime int64) (loginVo *LoginVo, statesCode string, bindKey string, err error) {

	return
}

func createLoginVo(loginUser interface{}, effectiveTime int64) LoginVo {
	return LoginVo{}
}
