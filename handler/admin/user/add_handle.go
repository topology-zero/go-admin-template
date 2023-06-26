package user

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/user"
	"go-admin-template/svc"
	userType "go-admin-template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加用户
func AddHandle(c *gin.Context) {
	var req userType.UserAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := user.Add(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, nil, err)
}
