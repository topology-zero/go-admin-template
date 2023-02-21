package auth

import (
	"github.com/pkg/errors"
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/auth"
	"go-admin-template/svc"
	authType "go-admin-template/types/admin/auth"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加权限
func AddHandle(c *gin.Context) {
	var req authType.AuthAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.IsMenu == 0 && (len(req.Api) == 0 || len(req.Action) == 0) {
		response.HandleResponse(c, nil, errors.New("非菜单权限必填操作方法和接口"))
		return
	}

	err := auth.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
