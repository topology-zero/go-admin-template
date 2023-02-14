package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"
	"admin_template/svc"

	"github.com/gin-gonic/gin"
)

// AuthHandle 获取所有权限
func AuthHandle(c *gin.Context) {
	resp, err := base.Auth(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
