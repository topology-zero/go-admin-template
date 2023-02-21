package base

import (
	"github.com/pkg/errors"
	"go-admin-template/pkg/util"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/base"
	"golang.org/x/crypto/bcrypt"
)

// ChangeSelfPwd 修改自己的密码
func ChangeSelfPwd(uid int, req *base.ChangeSelfPwdRequest, ctx *svc.ServiceContext) error {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.ID.Eq(uid)).First()
	err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.OldPassword))
	if err != nil {
		return errors.New("输入的原密码错误")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.Log.Errorf("%+v", err)
		return errors.New("系统错误")
	}

	_, err = userModel.Where(userModel.ID.Eq(uid)).Update(userModel.Password, string(password))
	return util.WarpDbError(err)
}
