package role

import (
	"strconv"

	"go-admin-template/config"
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Del 删除角色
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	if req.ID == config.SuperAdminRoleID {
		return errors.New("无法修改超级管理员角色")
	}

	adminUserModel := query.AdminUserModel

	user, _ := adminUserModel.WithContext(ctx).Where(adminUserModel.RoleID.Eq(req.ID)).First()
	if user != nil {
		return errors.New("当前角色正在使用,无法删除")
	}

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.WithContext(ctx).Where(tx.AdminRoleModel.ID.Eq(req.ID)).Delete()
		if err != nil {
			return err
		}

		_, err = tx.AdminCasbinRuleModel.WithContext(ctx).
			Where(
				tx.AdminCasbinRuleModel.Ptype.Eq("p"),
				tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.ID)),
			).
			Delete()
		return err
	})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
