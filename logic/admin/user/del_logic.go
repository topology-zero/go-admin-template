package user

import (
	"go-admin-template/config"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Del 删除用户
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	if req.ID == config.SuperAdminID {
		return errors.New("无法删除超级管理员")
	}
	adminUserModel := query.AdminUserModel

	_, err := adminUserModel.WithContext(ctx).
		Where(adminUserModel.ID.Eq(req.ID)).
		Delete()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
