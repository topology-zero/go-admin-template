package user

import (
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/user"
	"github.com/jinzhu/copier"
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
