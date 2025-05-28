package jwt

import (
	"time"

	"go-admin-template/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	AuthorizationName = "Authorization"
	UserInfo          = "userInfo"
)

type Claims struct {
	Id     int    `json:"id"`
	RoleId int    `json:"roleId"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

// MakeToken 生成 jwt 令牌
func MakeToken(id, roleId int, phone string) (string, error) {
	// 过期时间
	expTime := time.Now().Add(time.Duration(config.JwtConf.Expire) * time.Hour)
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Id:     id,
		RoleId: roleId,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			Subject:   config.ServerConf.Name,
		},
	})
	signedString, err := tokenClaim.SignedString([]byte(config.JwtConf.Secret))
	if err != nil {
		logrus.Errorf("生成jwt出错 : %+v", errors.WithStack(err))
	}
	return signedString, err
}

// ParseToken 解析 jwt 令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtConf.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	logrus.Errorf("解析jwt出错 : %+v", errors.WithStack(err))
	return nil, err
}

// GetToken 各种方法获取 token
func GetToken(c *gin.Context) (string, error) {
	if token := c.GetHeader(AuthorizationName); token != "" {
		return token, nil
	}

	if token, _ := c.Cookie(AuthorizationName); token != "" {
		return token, nil
	}

	if token := c.Query(AuthorizationName); token != "" {
		return token, nil
	}
	
	return "", errors.New("没有找到" + AuthorizationName)
}

// GetAndParseToken 对 GetToken 和 ParseToken 的封装
func GetAndParseToken(c *gin.Context) (*Claims, error) {
	token, err := GetToken(c)
	if err != nil {
		return nil, err
	}
	return ParseToken(token)
}
