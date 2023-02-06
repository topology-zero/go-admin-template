package user

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/user"
	userType "admin_template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加用户
func AddHandle(c *gin.Context) {
	var req userType.UserAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := user.Add(&req)
	response.HandleResponse(c, nil, err)
}
