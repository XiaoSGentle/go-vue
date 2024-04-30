package xtoken

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
	"xcore/common/xerror"
)

var ContextKey string = "BIND_CONTEXT_CUSTOM_CLAIMS_KEY"

// CustomClaims 自定义jwt的声明字段信息+标准字段
type CustomClaims struct {
	Uid   int64    `json:"uid"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateJwtToken
// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GenerateJwtToken(secretKey string, seconds int64, uid int64, roles []string) (token string, err error) {
	customClaims := CustomClaims{
		uid,
		roles,
		jwt.RegisteredClaims{
			Issuer:    "REACOOL_JWT_CLAIMS_ISSUER",
			Subject:   "REACOOL_JWT_CLAIMS_SUBJECT",
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+seconds, 0)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims).SignedString([]byte(secretKey))
	return
}

func ParseJwtToken(secretKey string, tokenString string) (claims *CustomClaims, err error) {
	parseClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return parseClaims.Claims.(*CustomClaims), err
	}
	if _claims, ok := parseClaims.Claims.(*CustomClaims); ok {
		return _claims, nil
	}
	return nil, xerror.NewErrCode(xerror.TOKEN_FORMAT_ERROR)
}

func RefreshToken(refreshToken string, refreshTokenSecretKey string, refreshTokenSecretKeyExpire int64, tokenSecretKey string, tokenSecretKeyExpire int64, uid int64, roles []string) (tokenString string, refreshTokenString string, err error) {
	_, err = jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenSecretKey), nil
	})
	tokenString, err = GenerateJwtToken(tokenSecretKey, tokenSecretKeyExpire, uid, roles)
	refreshTokenString, err = GenerateJwtToken(refreshTokenSecretKey, refreshTokenSecretKeyExpire, uid, roles)
	return
}

func GetClaimsByRequest(r *gin.Context, secretKey string) (customClaims *CustomClaims, err error) {
	authorizationHeader := r.GetHeader("Authorization")
	split := strings.Split(authorizationHeader, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		return nil, xerror.NewErrCode(xerror.TOKEN_FORMAT_ERROR)
	}
	customClaims, err = ParseJwtToken(secretKey, split[1])
	return
}
