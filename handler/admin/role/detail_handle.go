package role

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/role"
	"go-admin-template/svc"
	roleType "go-admin-template/types/admin/role"

	"github.com/gin-gonic/gin"
)

// DetailHandle 角色详情
func DetailHandle(c *gin.Context) {
	var req roleType.PathId
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := role.Detail(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
