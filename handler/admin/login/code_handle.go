package login

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/login"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// CodeHandle 获取验证码
func CodeHandle(c *gin.Context) {
	resp, err := login.Code(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
