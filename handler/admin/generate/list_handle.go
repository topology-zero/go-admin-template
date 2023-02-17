package generate

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/generate"
	"admin_template/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 表结构列表
func ListHandle(c *gin.Context) {
	resp, err := generate.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
