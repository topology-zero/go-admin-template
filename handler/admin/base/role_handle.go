package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// RoleHandle 获取所有角色
func RoleHandle(c *gin.Context) {
	resp, err := base.Role(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
