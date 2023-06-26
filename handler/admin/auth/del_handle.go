package auth

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/auth"
	"go-admin-template/svc"
	authType "go-admin-template/types/admin/auth"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除权限
func DelHandle(c *gin.Context) {
	var req authType.PathId
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := auth.Del(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
