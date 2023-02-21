package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// UserInfoHandle 获取用户信息
func UserInfoHandle(c *gin.Context) {
	res, err := base.UserInfo(svc.NewServiceContext(c))
	response.HandleResponse(c, res, err)
}
