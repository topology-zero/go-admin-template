package user

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/user"
	"go-admin-template/svc"
	userType "go-admin-template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// ListHandle 用户列表
func ListHandle(c *gin.Context) {
	var req userType.UserListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	resp, err := user.List(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
