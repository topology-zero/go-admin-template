package auth

import (
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/pkg/errors"
)

// Edit 编辑权限
func Edit(ctx *svc.ServiceContext, req *types.AdminAuthEditRequest) error {
	adminAuthModel := query.AdminAuthModel

	_, err := adminAuthModel.WithContext(ctx).Where(adminAuthModel.ID.Eq(req.ID)).UpdateColumnSimple(
		adminAuthModel.Pid.Value(req.Pid),
		adminAuthModel.Key.Value(req.Key),
		adminAuthModel.Name.Value(req.Name),
		adminAuthModel.IsMenu.Value(req.IsMenu),
		adminAuthModel.API.Value(req.API),
		adminAuthModel.Action.Value(req.Action),
	)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
