package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"
	"admin_template/svc"

	"github.com/gin-gonic/gin"
)

// RoleHandle 获取所有角色
func RoleHandle(c *gin.Context) {
	resp, err := base.Role(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
