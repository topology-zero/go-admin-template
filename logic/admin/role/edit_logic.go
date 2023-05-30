package role

import (
	"encoding/json"
	"strconv"

	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/role"

	"github.com/pkg/errors"
)

// Edit 编辑角色
func Edit(req *role.RoleEditRequest, ctx *svc.ServiceContext) error {
	if req.Id == AdminRoleId {
		return errors.New("无法修改超级管理员角色")
	}

	roleModel := query.AdminRoleModel
	roleInfo, _ := roleModel.WithContext(ctx).Where(roleModel.Name.Eq(req.Name), roleModel.ID.Neq(req.Id)).First()
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

	err = query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.WithContext(ctx).Where(tx.AdminRoleModel.ID.Eq(req.Id)).UpdateSimple(
			tx.AdminRoleModel.Name.Value(req.Name),
			tx.AdminRoleModel.Auth.Value(authStr),
		)
		if err != nil {
			return err
		}

		// 删除casbin表
		_, err = tx.AdminCasbinRuleModel.WithContext(ctx).Where(tx.AdminCasbinRuleModel.Ptype.Eq("p"), tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.Id))).Delete()
		if err != nil {
			return err
		}

		// 重新加入casbin表
		var rules []*model.AdminCasbinRuleModel
		authModels, _ := tx.AdminAuthModel.WithContext(ctx).Where(tx.AdminAuthModel.ID.In(req.Auth...)).Find()
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

		return tx.AdminCasbinRuleModel.WithContext(ctx).CreateInBatches(rules, 100)
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
