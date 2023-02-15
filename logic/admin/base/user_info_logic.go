package base

import (
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/base"
)

// UserInfo 获取用户信息
func UserInfo(id, roleId int, ctx *svc.ServiceContext) (resp base.UserInfoResponse, err error) {
	userModel := query.AdminUserModel
	resp = userModel.GetUserInfo(id)
	if roleId == 1 {
		resp.Authkeys = "*"
	}
	return
}