package role

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/role"
	roleType "admin_template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除角色
func DelHandle(c *gin.Context) {
	var req roleType.RoleDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Del(&req)
	response.HandleResponse(c, nil, err)
}
