package auth

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types"
)

// List 权限列表
func List(ctx *svc.ServiceContext) (resp []types.AuthListResponse, err error) {
	authModel := query.AdminAuthModel
	auths, _ := authModel.WithContext(ctx).Find()
	resp = authTree(0, auths)
	return
}

func authTree(pid int, auths []*model.AdminAuthModel) []types.AuthListResponse {
	var data []types.AuthListResponse
	for _, v := range auths {
		if v.Pid != pid {
			continue
		}
		data = append(data, types.AuthListResponse{
			ID:       v.ID,
			Pid:      v.Pid,
			Name:     v.Name,
			IsMenu:   v.IsMenu,
			Key:      v.Key,
			API:      v.API,
			Action:   v.Action,
			Children: authTree(v.ID, auths),
		})
	}
	return data
}
