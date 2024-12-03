package user

import (
	"go-admin-template/config"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gen/field"
)

// Edit 编辑用户
func Edit(ctx *svc.ServiceContext, req *types.AdminUserEditRequest) error {
	if req.ID == config.SuperAdminID && req.RoleID != config.SuperAdminRoleID {
		return errors.New("无法设置超级管理员的角色")
	}

	adminUserModel := query.AdminUserModel

	userInfo, _ := adminUserModel.WithContext(ctx).
		Where(adminUserModel.Username.Eq(req.Username), adminUserModel.ID.Neq(req.ID)).
		First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = adminUserModel.WithContext(ctx).
		Where(adminUserModel.Phone.Eq(req.Phone), adminUserModel.ID.Neq(req.ID)).
		First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	updates := []field.AssignExpr{
		adminUserModel.Username.Value(req.Username),
		adminUserModel.Realname.Value(req.Realname),
		adminUserModel.Phone.Value(req.Phone),
		adminUserModel.RoleID.Value(req.RoleID),
		adminUserModel.Status.Value(req.Status),
	}

	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.Log.Errorf("%+v", err)
			return errors.New("系统错误")
		}
		updates = append(updates, adminUserModel.Password.Value(string(password)))
	}

	_, err := adminUserModel.WithContext(ctx).
		Where(adminUserModel.ID.Eq(req.ID)).
		UpdateColumnSimple(updates...)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
