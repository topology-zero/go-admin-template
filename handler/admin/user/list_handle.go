package user

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/user"
	"admin_template/svc"
	userType "admin_template/types/admin/user"

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

	resp, err := user.List(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
