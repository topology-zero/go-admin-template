package role

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/role"
	"go-admin-template/svc"
	roleType "go-admin-template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加角色
func AddHandle(c *gin.Context) {
	var req roleType.RoleAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
