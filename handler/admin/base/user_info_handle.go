package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/pkg/jwt"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// UserInfoHandle 获取用户信息
func UserInfoHandle(c *gin.Context) {
	user, _ := c.Get("userInfo")
	claims := user.(*jwt.Claims)
	res, err := base.UserInfo(claims.Id, claims.RoleId, svc.NewServiceContext(c))
	response.HandleResponse(c, res, err)
}
