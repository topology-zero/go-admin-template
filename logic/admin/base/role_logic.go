package base

import (
	"github.com/jinzhu/copier"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/base"
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
