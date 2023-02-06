package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"
	"admin_template/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// UserInfoHandle 获取用户信息
func UserInfoHandle(c *gin.Context) {
	user, _ := c.Get("userInfo")
	claims := user.(*jwt.Claims)
	res, err := base.UserInfo(claims.Id, claims.RoleId)
	response.HandleResponse(c, res, err)
}
