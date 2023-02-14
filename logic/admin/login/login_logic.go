package login

import (
	"strings"

	"admin_template/config"
	"admin_template/pkg/jwt"
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/login"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Login 登录
func Login(req *login.LoginRequest, ctx *svc.ServiceContext) (resp login.LoginResponse, err error) {
	req.Code = strings.ToLower(req.Code)
	match := config.Captcha.Verify(req.CodeId, req.Code, true)
	if !match {
		err = errors.New("验证码错误")
		return
	}

	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.Username.Eq(req.Username)).First()
	if userInfo == nil {
		err = errors.New("用户名或密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	if userInfo.Status != 1 {
		err = errors.New("该用户已被管理员停用")
		return
	}

	token, err := jwt.MakeToken(userInfo.ID, userInfo.RoleID, userInfo.Phone)
	if err != nil {
		err = errors.New("系统错误")
		return
	}
	resp.Jwt = token
	return
}
