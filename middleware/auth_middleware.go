package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-admin-template/internal/response"
	"go-admin-template/model"
	"go-admin-template/pkg/jwt"
)

func AuthMiddleware(c *gin.Context) {
	user, _ := c.Get("userInfo")
	claims := user.(*jwt.Claims)
	roleStr := "role:" + strconv.Itoa(claims.RoleId)
	ok, _ := model.Enforcer.Enforce(roleStr, c.FullPath(), strings.ToLower(c.Request.Method))
	if !ok {
		logrus.Warning("没有权限")
		response.HandleAbortResponse(c, "没有权限", 403)
		return
	}
	c.Next()
}
