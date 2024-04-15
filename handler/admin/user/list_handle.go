package user

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/user"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/gin-gonic/gin"
)

// ListHandle 用户列表
func ListHandle(c *gin.Context) {
	var req types.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := user.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
