package user

import (
	"github.com/pkg/errors"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"
)

// Del 删除用户
func Del(req *user.UserDeleteRequest, ctx *svc.ServiceContext) error {
	if req.Id == SuperAdminID {
		return errors.New("无法删除超级管理员")
	}
	userModel := query.AdminUserModel
	_, err := userModel.Where(userModel.ID.Eq(req.Id)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
