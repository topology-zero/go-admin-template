package login

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/login"
	"admin_template/svc"

	"github.com/gin-gonic/gin"
)

// CodeHandle 获取验证码
func CodeHandle(c *gin.Context) {
	resp, err := login.Code(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
