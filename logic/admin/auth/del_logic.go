package auth

import (
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/types/admin/auth"
	"github.com/pkg/errors"
)

// Del 删除权限
func Del(req *auth.AuthDeleteRequest) error {
	authModel := query.AuthModel
	auths, _ := authModel.Where(authModel.Pid.Eq(req.Id)).First()
	if auths != nil {
		return errors.New("无法删除该权限,请先删除子权限")
	}

	roleModel := query.RoleModel
	role := roleModel.GetRoleByAuthId(req.Id)
	if role != nil {
		return errors.New("无法删除该权限,该权限已被使用")
	}

	_, err := authModel.Unscoped().Where(authModel.ID.Eq(req.Id)).Delete()
	return util.WarpDbError(err)
}
