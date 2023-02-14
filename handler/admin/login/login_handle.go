package login

import (
	"admin_template/config"
	"admin_template/internal/response"
	"admin_template/logic/admin/login"
	"admin_template/pkg/jwt"
	"admin_template/svc"
	loginType "admin_template/types/admin/login"

	"github.com/gin-gonic/gin"
)

// LoginHandle 登录
func LoginHandle(c *gin.Context) {
	var req loginType.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := login.Login(&req, svc.NewServiceContext(c))
	if err == nil {
		c.SetCookie(jwt.JwtName, resp.Jwt, config.JwtConf.Expire*3600, "/", "127.0.0.1", true, true)
		c.SetCookie(jwt.JwtName, resp.Jwt, config.JwtConf.Expire*3600, "/", "localhost", true, true)
	}
	response.HandleResponse(c, resp, err)
}
