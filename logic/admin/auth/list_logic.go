package auth

import (
	"admin_template/model"
	"admin_template/query"
	"admin_template/svc"
	"admin_template/types/admin/auth"
)

func List(ctx *svc.ServiceContext) (resp auth.AuthListResponse, err error) {
	authModel := query.AuthModel
	auths, _ := authModel.Find()
	resp.Tree = authTree(0, auths)
	return
}

func authTree(pid int, auths []*model.AuthModel) []auth.Auth {
	var data []auth.Auth
	for _, v := range auths {
		if v.Pid != pid {
			continue
		}
		data = append(data, auth.Auth{
			Id:     v.ID,
			Pid:    v.Pid,
			Name:   v.Name,
			IsMenu: v.IsMenu,
			Key:    v.Key,
			Api:    v.API,
			Action: v.Action,
			Child:  authTree(v.ID, auths),
		})
	}
	return data
}
