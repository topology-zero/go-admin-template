package auth

import (
	"admin_template/model"
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/types/admin/auth"
	"github.com/jinzhu/copier"
)

// Add 添加权限
func Add(req *auth.AuthAddRequest) error {
	authModel := query.AuthModel

	var saveAuth model.AuthModel
	copier.Copy(&saveAuth, &req)
	saveAuth.API = req.Api

	return util.WarpDbError(authModel.Create(&saveAuth))
}
