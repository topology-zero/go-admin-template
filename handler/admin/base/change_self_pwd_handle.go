package base

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/base"
	"admin_template/pkg/jwt"
	baseType "admin_template/types/admin/base"

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
	err := base.ChangeSelfPwd(claims.Id, &req)
	response.HandleResponse(c, nil, err)
}
