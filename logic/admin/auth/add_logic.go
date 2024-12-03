package auth

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Add 添加权限
func Add(ctx *svc.ServiceContext, req *types.AdminAuthAddRequest) error {
	authModel := query.AdminAuthModel

	err := authModel.WithContext(ctx).Create(&model.AdminAuthModel{
		Pid:    req.Pid,
		Name:   req.Name,
		Key:    req.Key,
		IsMenu: req.IsMenu,
		API:    req.API,
		Action: req.Action,
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
	}
	return err
}
