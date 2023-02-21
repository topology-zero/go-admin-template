package login

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/login"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// LogoutHandle 退出登录
func LogoutHandle(c *gin.Context) {
	err := login.Logout(svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
