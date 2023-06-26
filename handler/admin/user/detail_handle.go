package user

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/user"
	"go-admin-template/svc"
	userType "go-admin-template/types/admin/user"

	"github.com/gin-gonic/gin"
)

// DetailHandle 用户详情
func DetailHandle(c *gin.Context) {
	var req userType.PathId
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := user.Detail(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
