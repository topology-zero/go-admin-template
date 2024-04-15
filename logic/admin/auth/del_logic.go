package auth

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Del 删除权限
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	authModel := query.AdminAuthModel
	auths, _ := authModel.WithContext(ctx).Where(authModel.Pid.Eq(req.ID)).First()
	if auths != nil {
		return errors.New("无法删除该权限,请先删除子权限")
	}

	roleModel := query.AdminRoleModel
	role := roleModel.WithContext(ctx).GetRoleByAuthId(req.ID)
	if role.ID != 0 {
		return errors.New("无法删除该权限,该权限已被使用")
	}

	_, err := authModel.WithContext(ctx).Unscoped().Where(authModel.ID.Eq(req.ID)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
