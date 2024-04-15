package role

import (
	"encoding/json"
	"strconv"

	"go-admin-template/config"
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Edit 编辑角色
func Edit(ctx *svc.ServiceContext, req *types.RoleEditRequest) error {
	if req.ID == config.SuperAdminRoleID {
		return errors.New("无法修改超级管理员角色")
	}

	roleModel := query.AdminRoleModel
	roleInfo, _ := roleModel.WithContext(ctx).
		Where(
			roleModel.Name.Eq(req.Name),
			roleModel.ID.Neq(req.ID),
		).
		First()
	if roleInfo != nil {
		return errors.New("该角色已存在")
	}

	marshal, _ := json.Marshal(req.Auth)

	authStr := string(marshal)

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.WithContext(ctx).Where(tx.AdminRoleModel.ID.Eq(req.ID)).UpdateSimple(
			tx.AdminRoleModel.Name.Value(req.Name),
			tx.AdminRoleModel.Auth.Value(authStr),
		)
		if err != nil {
			return err
		}

		// 删除casbin表
		_, err = tx.AdminCasbinRuleModel.WithContext(ctx).
			Where(
				tx.AdminCasbinRuleModel.Ptype.Eq("p"),
				tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.ID)),
			).
			Delete()
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
				V0:    "role:" + strconv.Itoa(req.ID),
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
