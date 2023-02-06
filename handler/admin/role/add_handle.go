package role

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/role"
	roleType "admin_template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加角色
func AddHandle(c *gin.Context) {
	var req roleType.RoleAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Add(&req)
	response.HandleResponse(c, nil, err)
}
