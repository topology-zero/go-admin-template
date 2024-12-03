package role

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// List 角色列表
func List(ctx *svc.ServiceContext, req *types.AdminRoleListRequest) (resp types.AdminRoleListResponse, err error) {
	adminRoleModel := query.AdminRoleModel

	count, _ := adminRoleModel.WithContext(ctx).
		ScanByPage(&resp.Data, (req.Page-1)*req.PageSize, req.PageSize)

	resp.Total = int(count)
	return
}
