package user

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/user"
	userType "admin_template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// DetailHandle 用户详情
func DetailHandle(c *gin.Context) {
	var req userType.UserDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := user.Detail(&req)
	response.HandleResponse(c, resp, err)
}
