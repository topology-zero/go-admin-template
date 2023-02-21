package role

import (
	"github.com/jinzhu/copier"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/role"
)

// List 角色列表
func List(req *role.RoleListRequest, ctx *svc.ServiceContext) (resp role.RoleListResponse, err error) {
	roleModel := query.AdminRoleModel
	data, count, _ := roleModel.FindByPage((req.Page-1)*req.PageSize, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = int(count)
	copier.Copy(&resp.Data, &data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}
