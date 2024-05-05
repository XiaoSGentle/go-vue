package sys_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xencrypt"
	"xcore/common/xerror"
	xtoken "xcore/common/xtoken/jwt"
	"xcore/core/xvariable"
)

type IAuthService interface {
	UserLogin(c *gin.Context, param *LoginParam) (loginVo LoginVo, err error)
	GetUserById(c *gin.Context, uid int32) (user *model.SysUser, err error)
}

func NewAuthService(db *gorm.DB) IAuthService {
	return &AuthService{db: db, query: query.Use(db)}
}

type AuthService struct {
	db    *gorm.DB
	query *query.Query
}

func (s AuthService) GetUserById(c *gin.Context, uid int32) (user *model.SysUser, err error) {
	userQuery := s.query.SysUser
	user, err = userQuery.WithContext(c).Where(userQuery.ID.Eq(uid)).First()
	return
}

func (s AuthService) UserLogin(c *gin.Context, param *LoginParam) (loginVo LoginVo, err error) {
	userQuery := s.query.SysUser
	jwtTokenSignKey := xvariable.GlobalYmlConfig.GetString("Token.JwtTokenSignKey")
	jwtTokenCreatedExpireAt := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	maxLoginFailTimes := xvariable.GlobalYmlConfig.GetInt64("LoginPolicy.MaxLoginFailTimes")
	jwtTokenRefreshAllowSec := xvariable.GlobalYmlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")
	count, err := userQuery.Where(userQuery.Username.Eq(param.UserName)).Count()
	if err != nil || count < 1 {
		return LoginVo{}, xerror.NewErrCode(xerror.USER_NOT_EXIST_ERROR)
	}
	userInfoInSql, err := userQuery.Where(userQuery.Username.Eq(param.UserName)).First()
	if userInfoInSql == nil {
		return LoginVo{}, xerror.NewErrCode(xerror.USER_NOT_EXIST_ERROR)
	}
	if userInfoInSql.Password == xencrypt.Base64Md5(param.Password) {

		jwtToken, getTokenErr := xtoken.GenerateJwtToken(jwtTokenSignKey, jwtTokenCreatedExpireAt, &xtoken.ClaimsPayload{
			Uid:      userInfoInSql.ID,
			NickName: userInfoInSql.Nickname,
			Roles:    strings.Split(userInfoInSql.Roles, ","),
		})
		refreshToken, getTokenErr := xtoken.GenerateJwtToken(jwtTokenSignKey, jwtTokenRefreshAllowSec, &xtoken.ClaimsPayload{
			Uid:      userInfoInSql.ID,
			NickName: userInfoInSql.Nickname,
			Roles:    strings.Split(userInfoInSql.Roles, ","),
		})
		if getTokenErr != nil {
			return LoginVo{}, xerror.NewErrCode(xerror.TOKEN_GENERATE_ERROR)
		}
		_, _ = userQuery.Where(userQuery.Username.Eq(param.UserName)).Update(userQuery.LastOnlineTime, time.Now())

		return LoginVo{
			Token:        jwtToken,
			RefreshToken: refreshToken,
		}, nil
		// 密码不一致
	} else {
		split := strings.Split(userInfoInSql.LoginAttempts, "|")
		if split[0] == time.Now().Format(time.DateOnly) {
			// 这个是今天已有尝试登录
			toDayAttemptLoginTime, _ := strconv.ParseInt(split[1], 10, 32)
			toDayAttemptLoginTime++
			_, _ = userQuery.WithContext(c).Where(userQuery.Username.Eq(param.UserName)).Update(userQuery.LoginAttempts, time.Now().Format(time.DateOnly)+fmt.Sprintf("|%d", toDayAttemptLoginTime))

			var errMsg string
			if toDayAttemptLoginTime >= maxLoginFailTimes {
				errMsg = "您已被限制登录"
			} else {
				errMsg = fmt.Sprintf("密码错误,失败%d次后进入限制登录", maxLoginFailTimes-toDayAttemptLoginTime)
			}
			return LoginVo{}, xerror.NewErrCodeMsg(xerror.USER_PASSWORD_ERROR, errMsg)
		} else {
			// 未尝试登陆
			_, _ = userQuery.WithContext(c).Where(userQuery.Username.Eq(param.UserName)).Update(userQuery.LoginAttempts, time.Now().Format(time.DateOnly)+"|1")

			return LoginVo{}, xerror.NewErrCodeMsg(xerror.USER_PASSWORD_ERROR, fmt.Sprintf("密码错误,失败%d次后限制登录", maxLoginFailTimes-1))
		}
	}

}
