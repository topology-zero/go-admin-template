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
	userModel := query.AdminUserModel
	_, err := userModel.WithContext(ctx).Where(userModel.ID.Eq(req.ID)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
