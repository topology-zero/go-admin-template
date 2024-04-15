package middleware

import (
	"time"

	"go-admin-template/internal/response"
	"go-admin-template/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware(c *gin.Context) {
	token, err := jwt.GetToken(c)
	if err != nil {
		response.HandleAbortResponse(c, "鉴权失败，请重新登录", response.WithCode(401))
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		response.HandleAbortResponse(c, "鉴权失败，请重新登录", response.WithCode(401))
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		response.HandleAbortResponse(c, "鉴权失败，请重新登录", response.WithCode(401))
		return
	}

	// 保存进该请求的上下文
	c.Set("userInfo", claims)

	c.Next()
}
