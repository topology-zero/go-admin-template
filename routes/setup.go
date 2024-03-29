// Code generated by goctl. DO NOT EDIT.
package routes

import (
	adminAuth "go-admin-template/routes/admin/auth"
	adminBase "go-admin-template/routes/admin/base"
	adminLogin "go-admin-template/routes/admin/login"
	adminRole "go-admin-template/routes/admin/role"
	adminUser "go-admin-template/routes/admin/user"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) {
	adminLogin.RegisterAdminLoginRoute(e)
	adminUser.RegisterAdminUserRoute(e)
	adminRole.RegisterAdminRoleRoute(e)
	adminAuth.RegisterAdminAuthRoute(e)
	adminBase.RegisterAdminBaseRoute(e)
}
