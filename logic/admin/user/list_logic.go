package user

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// List 用户列表
func List(ctx *svc.ServiceContext, req *types.AdminUserListRequest) (resp types.AdminUserListResponse, err error) {
	adminUserModel := query.AdminUserModel
	adminRoleModel := query.AdminRoleModel

	count, err := adminUserModel.WithContext(ctx).
		LeftJoin(adminRoleModel, adminUserModel.RoleID.EqCol(adminRoleModel.ID)).
		Select(adminUserModel.ALL, adminRoleModel.Name.As("rolename")).
		ScanByPage(&resp.Data, (req.Page-1)*req.PageSize, req.PageSize)

	resp.Total = int(count)
	return
}
