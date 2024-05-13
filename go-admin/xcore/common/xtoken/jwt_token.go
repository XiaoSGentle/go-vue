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

type ClaimsPayload struct {
	Uid      int32    `json:"uid"`
	NickName string   `json:"nickName"`
	Roles    []string `json:"roles"`
}

// CustomClaims 自定义jwt的声明字段信息+标准字段
type CustomClaims struct {
	ClaimsPayload
	jwt.RegisteredClaims
}

// GenerateJwtToken
// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GenerateJwtToken(secretKey string, seconds int64, claimsPayload *ClaimsPayload) (token string, err error) {
	var customClaims CustomClaims
	customClaims.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "REACOOL_JWT_CLAIMS_ISSUER",
		Subject:   "REACOOL_JWT_CLAIMS_SUBJECT",
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+seconds, 0)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	}
	customClaims.ClaimsPayload = *claimsPayload
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims).SignedString([]byte(secretKey))
	return
}

// ParseJwtToken  解析Token
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

// RefreshToken  刷新Token
func RefreshToken(refreshToken string, refreshTokenSecretKey string, refreshTokenSecretKeyExpire int64, tokenSecretKey string, tokenSecretKeyExpire int64, payload *ClaimsPayload) (tokenString string, refreshTokenString string, err error) {
	_, err = jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenSecretKey), nil
	})
	tokenString, err = GenerateJwtToken(tokenSecretKey, tokenSecretKeyExpire, payload)
	refreshTokenString, err = GenerateJwtToken(refreshTokenSecretKey, refreshTokenSecretKeyExpire, payload)
	return
}

// GetPayloadByRequest 从gin.Context获取PayLoad
func GetPayloadByRequest(r *gin.Context, secretKey string) (claimsPayload *ClaimsPayload, err error) {
	authorizationHeader := r.GetHeader("Authorization")
	split := strings.Split(authorizationHeader, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		return nil, xerror.NewErrCode(xerror.TOKEN_FORMAT_ERROR)
	}
	customClaims, err := ParseJwtToken(secretKey, split[1])
	return &customClaims.ClaimsPayload, nil
}
func GetClaimsByRequest(r *gin.Context, secretKey string) (claimsPayload *CustomClaims, err error) {
	authorizationHeader := r.GetHeader("Authorization")
	split := strings.Split(authorizationHeader, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		return nil, xerror.NewErrCode(xerror.TOKEN_FORMAT_ERROR)
	}
	customClaims, err := ParseJwtToken(secretKey, split[1])
	if err != nil {
		return customClaims, err
	}
	return customClaims, nil
}

func GetBindCustomPayload(c *gin.Context) *ClaimsPayload {
	claimsPayload, exists := c.Get(ContextKey)
	if exists {
		payload := claimsPayload.(ClaimsPayload)
		return &ClaimsPayload{
			Uid:      payload.Uid,
			NickName: payload.NickName,
			Roles:    payload.Roles,
		}
	} else {
		return &ClaimsPayload{
			Uid:      -1,
			NickName: "NoAuthMethod",
			Roles:    []string{},
		}
	}
}
