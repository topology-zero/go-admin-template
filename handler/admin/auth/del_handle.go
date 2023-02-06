package auth

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/auth"
	authType "admin_template/types/admin/auth"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除权限
func DelHandle(c *gin.Context) {
	var req authType.AuthDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := auth.Del(&req)
	response.HandleResponse(c, nil, err)
}
