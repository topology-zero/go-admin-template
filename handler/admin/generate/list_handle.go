package generate

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/generate"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 表结构列表
func ListHandle(c *gin.Context) {
	resp, err := generate.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
