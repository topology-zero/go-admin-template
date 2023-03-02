package role

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/role"
)

// Add 添加角色
func Add(req *role.RoleAddRequest, ctx *svc.ServiceContext) error {
	roleModel := query.AdminRoleModel
	roleInfo, _ := roleModel.Where(roleModel.Name.Eq(req.Name)).First()
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
		saveRole := model.AdminRoleModel{
			Name: req.Name,
			Auth: authStr,
		}
		err = tx.AdminRoleModel.Create(&saveRole)
		if err != nil {
			return err
		}

		var rules []*model.AdminCasbinRuleModel

		authModels, _ := tx.AdminAuthModel.Where(tx.AdminAuthModel.ID.In(req.Auth...)).Find()
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

		return tx.AdminCasbinRuleModel.CreateInBatches(rules, 100)
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
