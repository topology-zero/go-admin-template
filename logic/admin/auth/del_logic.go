package auth

import (
	"github.com/pkg/errors"
	"go-admin-template/pkg/util"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/auth"
)

// Del 删除权限
func Del(req *auth.AuthDeleteRequest, ctx *svc.ServiceContext) error {
	authModel := query.AdminAuthModel
	auths, _ := authModel.Where(authModel.Pid.Eq(req.Id)).First()
	if auths != nil {
		return errors.New("无法删除该权限,请先删除子权限")
	}

	roleModel := query.AdminRoleModel
	role := roleModel.GetRoleByAuthId(req.Id)
	if role.ID != 0 {
		return errors.New("无法删除该权限,该权限已被使用")
	}

	_, err := authModel.Unscoped().Where(authModel.ID.Eq(req.Id)).Delete()
	return util.WarpDbError(err)
}
