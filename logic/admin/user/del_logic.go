package user

import (
	"github.com/pkg/errors"
	"go-admin-template/pkg/util"
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
	return util.WarpDbError(err)
}
