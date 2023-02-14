package user

import (
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/user"
	"github.com/pkg/errors"
)

// Del 删除用户
func Del(req *user.UserDeleteRequest, ctx *svc.ServiceContext) error {
	if req.Id == SuperAdminID {
		return errors.New("无法删除超级管理员")
	}
	userModel := query.AdminUserModel
	_, err := userModel.Where(userModel.ID.Eq(req.Id)).Delete()
	return util.WarpDbError(err)
}
