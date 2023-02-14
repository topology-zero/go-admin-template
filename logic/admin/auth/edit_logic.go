package auth

import (
	"admin_template/pkg/util"
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/auth"
)

// Edit 编辑权限
func Edit(req *auth.AuthEditRequest, ctx *svc.ServiceContext) error {
	authModel := query.AuthModel
	_, err := authModel.Where(authModel.ID.Eq(req.Id)).UpdateSimple(
		authModel.Pid.Value(req.Pid),
		authModel.Key.Value(req.Key),
		authModel.Name.Value(req.Name),
		authModel.IsMenu.Value(req.IsMenu),
		authModel.API.Value(req.Api),
		authModel.Action.Value(req.Action),
	)
	return util.WarpDbError(err)
}
