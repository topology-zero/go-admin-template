package base

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/base"

	"github.com/jinzhu/copier"
)

// Role 获取所有角色
func Role(ctx *svc.ServiceContext) (resp base.BaseRoleResponse, err error) {
	roleModel := query.AdminRoleModel
	data, _ := roleModel.WithContext(ctx).Find()
	copier.Copy(&resp.Data, data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}
