package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// AuthHandle 获取所有权限
func AuthHandle(c *gin.Context) {
	resp, err := base.Auth(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
