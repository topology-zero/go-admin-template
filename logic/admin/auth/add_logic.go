package auth

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-admin-template/model"
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

	err := authModel.Create(&saveAuth)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
