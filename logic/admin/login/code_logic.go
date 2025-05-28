package login

import (
	"go-admin-template/config"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// Code 获取验证码
func Code(_ *svc.ServiceContext) (resp types.CodeResponse, err error) {
	resp.ID, resp.Image, _, err = config.Captcha.Generate()
	return
}
