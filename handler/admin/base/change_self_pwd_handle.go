package base

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/base"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
)

// ChangeSelfPwdHandle 修改自己的密码
func ChangeSelfPwdHandle(c *gin.Context) {
	var req types.ChangeSelfPwdRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	err := base.ChangeSelfPwd(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
