package generate

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/generate"
	"admin_template/svc"
	generateType "admin_template/types/admin/generate"

	"github.com/gin-gonic/gin"
)

// GenerateHandle 生成前端文件
func GenerateHandle(c *gin.Context) {
	var req generateType.GenerateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := generate.Generate(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
