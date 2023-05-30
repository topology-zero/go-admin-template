package role

import (
	"encoding/json"

	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/role"

	"github.com/pkg/errors"
)

// Detail 角色详情
func Detail(req *role.PathId, ctx *svc.ServiceContext) (resp role.RoleDetailResponse, err error) {
	roleModel := query.AdminRoleModel
	authModel := query.AdminAuthModel
	roleInfo, _ := roleModel.WithContext(ctx).Where(roleModel.ID.Eq(req.Id)).First()

	var auths []int
	err = json.Unmarshal([]byte(roleInfo.Auth), &auths)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("JSON转换错误")
		return
	}

	var auth []int
	_ = authModel.WithContext(ctx).Select(authModel.ID).Where(authModel.IsMenu.Eq(0), authModel.ID.In(auths...)).Scan(&auth)
	resp.Id = roleInfo.ID
	resp.Name = roleInfo.Name
	resp.Auth = auth
	return
}
