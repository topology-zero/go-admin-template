package auth

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/auth"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 权限列表
func ListHandle(c *gin.Context) {
	resp, err := auth.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
