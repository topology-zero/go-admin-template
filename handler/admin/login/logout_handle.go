package login

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/login"

	"github.com/gin-gonic/gin"
)

// LogoutHandle 退出登录
func LogoutHandle(c *gin.Context) {
	err := login.Logout()
	response.HandleResponse(c, nil, err)
}
