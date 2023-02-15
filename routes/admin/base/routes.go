// Code generated by goctl. DO NOT EDIT.
package base

import (
	"admin_template/handler/admin/base"
	"admin_template/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminBaseRoute(e *gin.Engine) {
	g := e.Group("/base")
	g.Use(middleware.JwtMiddleware)
	g.GET("/role", base.RoleHandle)
	g.GET("/auth", base.AuthHandle)
	g.GET("/userInfo", base.UserInfoHandle)
	g.POST("/changeSelfPwd", base.ChangeSelfPwdHandle)

}