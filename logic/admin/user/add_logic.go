package user

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Add 添加用户
func Add(ctx *svc.ServiceContext, req *user.UserAddRequest) error {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.WithContext(ctx).Where(userModel.Username.Eq(req.Username)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = userModel.WithContext(ctx).Where(userModel.Phone.Eq(req.Phone)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(password)
	if err != nil {
		ctx.Log.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	var u model.AdminUserModel
	copier.Copy(&u, &req)
	u.RoleID = req.RoleId

	err = userModel.WithContext(ctx).Create(&u)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
