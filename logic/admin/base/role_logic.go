package base

import (
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/base"
	"github.com/jinzhu/copier"
)

// Role 获取所有角色
func Role(ctx *svc.ServiceContext) (resp base.BaseRoleResponse, err error) {
	roleModel := query.AdminRoleModel
	data, _ := roleModel.Find()
	copier.Copy(&resp.Data, data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}
