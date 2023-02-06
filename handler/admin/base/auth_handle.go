package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"

	"github.com/gin-gonic/gin"
)

// AuthHandle 获取所有权限
func AuthHandle(c *gin.Context) {
	resp, err := base.Auth()
	response.HandleResponse(c, resp, err)
}
