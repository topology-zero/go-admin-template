package login

import (
	"go-admin-template/config"
	"go-admin-template/svc"
	"go-admin-template/types/admin/login"
)

// Code 获取验证码
func Code(ctx *svc.ServiceContext) (resp login.CodeResponse, err error) {
	id, image, err := config.Captcha.Generate()
	if err != nil {
		return
	}

	resp.Id = id
	resp.Image = image
	return
}
