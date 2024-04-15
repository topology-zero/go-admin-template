package role

import (
	"encoding/json"

	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// Detail 角色详情
func Detail(ctx *svc.ServiceContext, req *types.PathID) (resp types.RoleDetailResponse, err error) {
	roleModel := query.AdminRoleModel
	authModel := query.AdminAuthModel
	roleInfo, _ := roleModel.WithContext(ctx).Where(roleModel.ID.Eq(req.ID)).First()

	var auths []int
	_ = json.Unmarshal([]byte(roleInfo.Auth), &auths)

	var auth []int
	_ = authModel.WithContext(ctx).
		Select(authModel.ID).
		Where(
			authModel.IsMenu.Eq(0),
			authModel.ID.In(auths...),
		).
		Scan(&auth)
	resp.ID = roleInfo.ID
	resp.Name = roleInfo.Name
	resp.Auth = auth
	return
}
