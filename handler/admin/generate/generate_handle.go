package generate

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/generate"
	"go-admin-template/svc"
	generateType "go-admin-template/types/admin/generate"

	"github.com/gin-gonic/gin"
)

// GenerateHandle 生成前端文件
func GenerateHandle(c *gin.Context) {
	var req generateType.GenerateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := generate.Generate(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
