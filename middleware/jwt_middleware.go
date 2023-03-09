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
		response.HandleAbortResponse(c, "鉴权失败，没有获取到有效的 token")
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		response.HandleAbortResponse(c, "鉴权失败，没有获取到有效的 token")
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		response.HandleAbortResponse(c, "鉴权失败，没有获取到有效的 token")
		return
	}

	// 保存进该请求的上下文
	c.Set("userInfo", claims)

	c.Next()
}
