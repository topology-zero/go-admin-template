package role

import (
	"strconv"

	"admin_template/model"
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/types/admin/role"
	"github.com/pkg/errors"
)

// Del 删除角色
func Del(req *role.RoleDeleteRequest) error {
	if req.Id == AdminRoleId {
		return errors.New("无法修改超级管理员角色")
	}

	userModel := query.AdminUserModel
	user, _ := userModel.Where(userModel.RoleID.Eq(req.Id)).First()
	if user != nil {
		return errors.New("当前角色正在使用,无法删除")
	}

	defer model.Enforcer.LoadPolicy()

	return util.WarpDbError(query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.RoleModel.Where(tx.RoleModel.ID.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}

		_, err = tx.CasbinRuleModel.Where(tx.CasbinRuleModel.Ptype.Eq("p"), tx.CasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.Id))).Delete()
		return err
	}))

}
