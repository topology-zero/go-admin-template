package role

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/role"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 角色列表
func ListHandle(c *gin.Context) {
	var req types.AdminRoleListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := role.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
