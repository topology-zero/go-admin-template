package role

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/role"
	"go-admin-template/svc"
	roleType "go-admin-template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑角色
func EditHandle(c *gin.Context) {
	var req roleType.RoleEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
