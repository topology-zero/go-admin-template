package role

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/role"

	"github.com/jinzhu/copier"
)

// List 角色列表
func List(ctx *svc.ServiceContext, req *role.RoleListRequest) (resp role.RoleListResponse, err error) {
	roleModel := query.AdminRoleModel
	data, count, _ := roleModel.WithContext(ctx).FindByPage((req.Page-1)*req.PageSize, req.PageSize)
	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}
