package user

import (
	"github.com/jinzhu/copier"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"
)

// Detail 用户详情
func Detail(req *user.UserDetailRequest, ctx *svc.ServiceContext) (resp user.UserDetailResponse, err error) {
	userModel := query.AdminUserModel
	userInfo, _ := userModel.Where(userModel.ID.Eq(req.Id)).First()
	copier.Copy(&resp, &userInfo)
	resp.Id = userInfo.ID
	resp.RoleId = userInfo.RoleID
	return
}
