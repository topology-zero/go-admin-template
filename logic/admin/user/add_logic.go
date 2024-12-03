package user

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Add 添加用户
func Add(ctx *svc.ServiceContext, req *types.AdminUserAddRequest) error {
	adminUserModel := query.AdminUserModel

	userInfo, _ := adminUserModel.WithContext(ctx).Where(adminUserModel.Username.Eq(req.Username)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = adminUserModel.WithContext(ctx).Where(adminUserModel.Phone.Eq(req.Phone)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Log.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	err = adminUserModel.WithContext(ctx).
		Create(&model.AdminUserModel{
			Username: req.Username,
			Realname: req.Realname,
			Password: string(password),
			Phone:    req.Phone,
			RoleID:   req.RoleID,
			Status:   req.Status,
		})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
