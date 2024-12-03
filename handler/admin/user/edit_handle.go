package user

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/user"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// EditHandle 编辑用户
func EditHandle(c *gin.Context) {
	var req types.AdminUserEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	pwdLen := len(req.Password)
	if pwdLen > 0 && pwdLen < 6 {
		response.HandleResponse(c, nil, errors.New("密码不得小于6位数"))
		return
	}

	err := user.Edit(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
