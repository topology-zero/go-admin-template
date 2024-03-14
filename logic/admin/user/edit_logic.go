package user

import (
	"go-admin-template/config"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gen/field"
)

// Edit 编辑用户
func Edit(ctx *svc.ServiceContext, req *user.UserEditRequest) error {
	if req.Id == config.SuperAdminID && req.RoleId != config.SuperAdminRoleID {
		return errors.New("无法设置超级管理员的角色")
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.Username.Eq(req.Username), userModel.ID.Neq(req.Id)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = userModel.WithContext(ctx).Where(userModel.Phone.Eq(req.Phone), userModel.ID.Neq(req.Id)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	updates := []field.AssignExpr{
		userModel.Username.Value(req.Username),
		userModel.Realname.Value(req.Realname),
		userModel.Phone.Value(req.Phone),
		userModel.RoleID.Value(req.RoleId),
		userModel.Status.Value(req.Status),
	}

	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.Log.Errorf("%+v", err)
			return errors.New("系统错误")
		}
		updates = append(updates, userModel.Password.Value(string(password)))
	}

	_, err := userModel.WithContext(ctx).Where(userModel.ID.Eq(req.Id)).UpdateColumnSimple(updates...)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
