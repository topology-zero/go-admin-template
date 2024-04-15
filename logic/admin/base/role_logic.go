package base

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// Role 获取所有角色
func Role(ctx *svc.ServiceContext) (resp []types.IDAndName, err error) {
	roleModel := query.AdminRoleModel
	_ = roleModel.WithContext(ctx).Select(roleModel.ID, roleModel.Name).Scan(&resp)
	return
}
