package common

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/common"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// UploadImageHandle 上传图片
func UploadImageHandle(c *gin.Context) {
	resp, err := common.UploadImage(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
