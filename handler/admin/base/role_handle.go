package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"

	"github.com/gin-gonic/gin"
)

// RoleHandle 获取所有角色
func RoleHandle(c *gin.Context) {
	resp, err := base.Role()
	response.HandleResponse(c, resp, err)
}
