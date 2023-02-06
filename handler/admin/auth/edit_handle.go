package auth

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/auth"
	authType "admin_template/types/admin/auth"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑权限
func EditHandle(c *gin.Context) {
	var req authType.AuthEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := auth.Edit(&req)
	response.HandleResponse(c, nil, err)
}
