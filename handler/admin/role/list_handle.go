package role

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/role"
	roleType "admin_template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// ListHandle 角色列表
func ListHandle(c *gin.Context) {
	var req roleType.RoleListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	resp, err := role.List(&req)
	response.HandleResponse(c, resp, err)
}
