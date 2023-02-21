package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/pkg/jwt"
	"go-admin-template/svc"
	baseType "go-admin-template/types/admin/base"

	"github.com/gin-gonic/gin"
)

// ChangeSelfPwdHandle 修改自己的密码
func ChangeSelfPwdHandle(c *gin.Context) {
	var req baseType.ChangeSelfPwdRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	userInfo, _ := c.Get("userInfo")
	claims := userInfo.(*jwt.Claims)
	err := base.ChangeSelfPwd(claims.Id, &req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
