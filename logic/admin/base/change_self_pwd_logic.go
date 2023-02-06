package base

import (
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/types/admin/base"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// ChangeSelfPwd 修改自己的密码
func ChangeSelfPwd(uid int, req *base.ChangeSelfPwdRequest) error {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.ID.Eq(uid)).First()
	err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.OldPassword))
	if err != nil {
		return errors.New("输入的原密码错误")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	_, err = userModel.Where(userModel.ID.Eq(uid)).Update(userModel.Password, string(password))
	return util.WarpDbError(err)
}
