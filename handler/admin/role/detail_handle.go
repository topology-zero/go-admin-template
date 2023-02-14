package role

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/role"
	"admin_template/svc"
	roleType "admin_template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// DetailHandle 角色详情
func DetailHandle(c *gin.Context) {
	var req roleType.RoleDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := role.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
