package auth

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/auth"

	"github.com/gin-gonic/gin"
)

// ListHandle 权限列表
func ListHandle(c *gin.Context) {
	resp, err := auth.List()
	response.HandleResponse(c, resp, err)
}
