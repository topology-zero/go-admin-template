package login

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/login"
	"admin_template/svc"

	"github.com/gin-gonic/gin"
)

// LogoutHandle 退出登录
func LogoutHandle(c *gin.Context) {
	err := login.Logout(svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
