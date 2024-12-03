package auth

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/auth"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑权限
func EditHandle(c *gin.Context) {
	var req types.AdminAuthEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := auth.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
