package user

import (
	"github.com/pkg/errors"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"
	"golang.org/x/crypto/bcrypt"
)

// Edit 编辑用户
func Edit(req *user.UserEditRequest, ctx *svc.ServiceContext) error {
	if req.Id == SuperAdminID && req.RoleId != SuperAdminRoleID {
		return errors.New("无法设置超级管理员的角色")
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.Username.Eq(req.Username), userModel.ID.Neq(req.Id)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = userModel.Where(userModel.Phone.Eq(req.Phone), userModel.ID.Neq(req.Id)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	updates := map[string]any{
		"username": req.Username,
		"realname": req.Realname,
		"phone":    req.Phone,
		"role_id":  req.RoleId,
		"status":   req.Status,
	}

	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.Log.Errorf("%+v", err)
			return errors.New("系统错误")
		}
		updates["password"] = string(password)
	}

	_, err := userModel.Where(userModel.ID.Eq(req.Id)).Updates(updates)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
