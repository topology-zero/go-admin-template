package role

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/jinzhu/copier"
)

// List 角色列表
func List(ctx *svc.ServiceContext, req *types.RoleListRequest) (resp types.RoleListResponse, err error) {
	roleModel := query.AdminRoleModel
	data, count, _ := roleModel.WithContext(ctx).FindByPage((req.Page-1)*req.PageSize, req.PageSize)
	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	return
}
