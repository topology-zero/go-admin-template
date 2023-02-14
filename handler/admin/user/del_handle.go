package user

import (
	"admin_template/internal/response"
	"admin_template/logic/admin/user"
	"admin_template/svc"
	userType "admin_template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除用户
func DelHandle(c *gin.Context) {
	var req userType.UserDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := user.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
