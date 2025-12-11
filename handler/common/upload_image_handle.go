package common

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/common"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
)

// UploadImageHandle 上传图片
func UploadImageHandle(c *gin.Context) {
	var req types.UploadImageRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := common.UploadImage(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
