package login

import (
	"admin_template/config"
	"admin_template/types/admin/login"
)

// Code 获取验证码
func Code() (resp login.CodeResponse, err error) {
	id, image, err := config.Captcha.Generate()
	if err != nil {
		return
	}

	resp.Id = id
	resp.Image = image
	return
}
