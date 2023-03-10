package auth

import (
	"go-admin-template/internal/response"
	"go-admin-template/logic/admin/auth"
	"go-admin-template/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle ๆ้ๅ่กจ
func ListHandle(c *gin.Context) {
	resp, err := auth.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
