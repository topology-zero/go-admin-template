package base

import (
	"go-admin-template/pkg/jwt"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// UserInfo 获取用户信息
func UserInfo(ctx *svc.ServiceContext) (resp types.UserInfoResponse, err error) {
	user, _ := ctx.GinContext.Get(jwt.UserInfo)
	claims := user.(*jwt.Claims)
	userModel := query.AdminUserModel
	resp = userModel.WithContext(ctx).GetUserInfo(claims.Id)
	if claims.RoleId == 1 {
		resp.Authkeys = "*"
	}
	return
}
