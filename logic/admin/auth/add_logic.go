package auth

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

// Add 添加权限
func Add(ctx *svc.ServiceContext, req *types.AuthAddRequest) error {
	authModel := query.AdminAuthModel

	var saveAuth model.AdminAuthModel
	copier.Copy(&saveAuth, &req)

	err := authModel.WithContext(ctx).Create(&saveAuth)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
