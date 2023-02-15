// Code generated by goctl. DO NOT EDIT.
package auth

import (
	"admin_template/handler/admin/auth"
	"admin_template/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminAuthRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware, middleware.AuthMiddleware)
	g.GET("/auth", auth.ListHandle)
	g.POST("/auth", auth.AddHandle)
	g.PUT("/auth/:id", auth.EditHandle)
	g.DELETE("/auth/:id", auth.DelHandle)

}