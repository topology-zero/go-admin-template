package role

import (
	"encoding/json"
	"strconv"

	"admin_template/model"
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/role"
	"github.com/pkg/errors"
)

// Edit 编辑角色
func Edit(req *role.RoleEditRequest, ctx *svc.ServiceContext) error {
	if req.Id == AdminRoleId {
		return errors.New("无法修改超级管理员角色")
	}

	roleModel := query.AdminRoleModel
	roleInfo, _ := roleModel.Where(roleModel.Name.Eq(req.Name), roleModel.ID.Neq(req.Id)).First()
	if roleInfo != nil {
		return errors.New("该角色已存在")
	}

	marshal, err := json.Marshal(req.Auth)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return errors.New("JSON转换错误")
	}
	authStr := string(marshal)

	defer model.Enforcer.LoadPolicy()

	return util.WarpDbError(query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.Where(tx.AdminRoleModel.ID.Eq(req.Id)).UpdateSimple(
			tx.AdminRoleModel.Name.Value(req.Name),
			tx.AdminRoleModel.Auth.Value(authStr),
		)
		if err != nil {
			return err
		}

		// 删除casbin表
		_, err = tx.AdminCasbinRuleModel.Where(tx.AdminCasbinRuleModel.Ptype.Eq("p"), tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.Id))).Delete()
		if err != nil {
			return err
		}

		// 重新加入casbin表
		var rules []*model.AdminCasbinRuleModel
		authModels, _ := tx.AdminAuthModel.Where(tx.AdminAuthModel.ID.In(req.Auth...)).Find()
		for _, a := range authModels {
			if a.IsMenu == 1 {
				continue
			}

			rules = append(rules, &model.AdminCasbinRuleModel{
				Ptype: "p",
				V0:    "role:" + strconv.Itoa(req.Id),
				V1:    a.API,
				V2:    a.Action,
			})
		}

		return tx.AdminCasbinRuleModel.CreateInBatches(rules, 100)
	}))
}
