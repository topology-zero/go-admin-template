package user

import (
	"admin_template/model"
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/types/admin/user"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Add 添加用户
func Add(req *user.UserAddRequest) error {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.Username.Eq(req.Username)).First()
	if userInfo != nil {
		return errors.New("该账号已被注册")
	}

	userInfo, _ = userModel.Where(userModel.Phone.Eq(req.Phone)).First()
	if userInfo != nil {
		return errors.New("该手机已被注册")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(password)
	if err != nil {
		logrus.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	var u model.AdminUserModel
	copier.Copy(&u, &req)
	u.RoleID = req.RoleId

	return util.WarpDbError(userModel.Create(&u))
}
