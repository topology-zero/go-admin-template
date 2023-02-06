package user

import (
	"admin_template/query"
	"admin_template/types/admin/user"

	"github.com/jinzhu/copier"
)

// List 用户列表
func List(req *user.UserListRequest) (resp user.UserListResponse, err error) {
	userModel := query.AdminUserModel
	roleModel := query.RoleModel
	var data []user.UserList

	count, err := userModel.LeftJoin(roleModel, userModel.RoleID.EqCol(roleModel.ID)).
		Select(userModel.ALL, roleModel.Name.As("rolename")).
		ScanByPage(&data, (req.Page-1)*req.PageSize, req.PageSize)

	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	return
}
