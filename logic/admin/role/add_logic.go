package role

import (
	"encoding/json"
	"strconv"

	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Add 添加角色
func Add(ctx *svc.ServiceContext, req *types.AdminRoleAddRequest) error {
	adminRoleModel := query.AdminRoleModel

	roleInfo, _ := adminRoleModel.WithContext(ctx).Where(adminRoleModel.Name.Eq(req.Name)).First()
	if roleInfo != nil {
		return errors.New("该角色已存在")
	}

	marshal, _ := json.Marshal(req.Auth)
	authStr := string(marshal)

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		saveRole := model.AdminRoleModel{
			Name: req.Name,
			Auth: authStr,
		}
		err := tx.AdminRoleModel.WithContext(ctx).Create(&saveRole)
		if err != nil {
			return err
		}

		var rules []*model.AdminCasbinRuleModel

		authModels, _ := tx.AdminAuthModel.WithContext(ctx).Where(tx.AdminAuthModel.ID.In(req.Auth...)).Find()
		for _, a := range authModels {
			if a.IsMenu == 1 {
				continue
			}

			rules = append(rules, &model.AdminCasbinRuleModel{
				Ptype: "p",
				V0:    "role:" + strconv.Itoa(saveRole.ID),
				V1:    a.API,
				V2:    a.Action,
			})
		}

		return tx.AdminCasbinRuleModel.WithContext(ctx).CreateInBatches(rules, 100)
	})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
