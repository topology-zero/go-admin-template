package auth

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Del 删除权限
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	adminAuthModel := query.AdminAuthModel
	adminRoleModel := query.AdminRoleModel

	auths, _ := adminAuthModel.WithContext(ctx).
		Where(adminAuthModel.Pid.Eq(req.ID)).
		First()
	if auths != nil {
		return errors.New("无法删除该权限,请先删除子权限")
	}

	role := adminRoleModel.WithContext(ctx).
		GetRoleByAuthId(req.ID)
	if role.ID != 0 {
		return errors.New("无法删除该权限,该权限已被使用")
	}

	_, err := adminAuthModel.WithContext(ctx).
		Unscoped().
		Where(adminAuthModel.ID.Eq(req.ID)).
		Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
	}
	return err
}
