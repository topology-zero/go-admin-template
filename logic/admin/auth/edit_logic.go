package auth

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Edit 编辑权限
func Edit(ctx *svc.ServiceContext, req *types.AuthEditRequest) error {
	authModel := query.AdminAuthModel
	_, err := authModel.WithContext(ctx).Where(authModel.ID.Eq(req.ID)).UpdateSimple(
		authModel.Pid.Value(req.Pid),
		authModel.Key.Value(req.Key),
		authModel.Name.Value(req.Name),
		authModel.IsMenu.Value(req.IsMenu),
		authModel.API.Value(req.API),
		authModel.Action.Value(req.Action),
	)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
