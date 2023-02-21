package auth

import (
	"github.com/jinzhu/copier"
	"go-admin-template/model"
	"go-admin-template/pkg/util"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/auth"
)

// Add 添加权限
func Add(req *auth.AuthAddRequest, ctx *svc.ServiceContext) error {
	authModel := query.AdminAuthModel

	var saveAuth model.AdminAuthModel
	copier.Copy(&saveAuth, &req)
	saveAuth.API = req.Api

	return util.WarpDbError(authModel.Create(&saveAuth))
}
