package user

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/user"

	"github.com/jinzhu/copier"
)

// List 用户列表
func List(ctx *svc.ServiceContext, req *user.UserListRequest) (resp user.UserListResponse, err error) {
	userModel := query.AdminUserModel
	roleModel := query.AdminRoleModel
	var data []user.UserList

	count, err := userModel.WithContext(ctx).LeftJoin(roleModel, userModel.RoleID.EqCol(roleModel.ID)).
		Select(userModel.ALL, roleModel.Name.As("rolename")).
		ScanByPage(&data, (req.Page-1)*req.PageSize, req.PageSize)

	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	return
}
