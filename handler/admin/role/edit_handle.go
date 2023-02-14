package role

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/role"
	"admin_template/svc"
	roleType "admin_template/types/admin/role"

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

	err := role.Edit(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
