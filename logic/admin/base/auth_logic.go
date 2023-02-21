package base

import (
	"go-admin-template/model"
	"go-admin-template/query"
	"go-admin-template/svc"
	"go-admin-template/types/admin/base"
)

func Auth(ctx *svc.ServiceContext) (resp base.BaseAuthResponse, err error) {
	authModel := query.AdminAuthModel
	auths, _ := authModel.Find()
	resp.Tree = authTree(0, auths)
	return
}

func authTree(pid int, auths []*model.AdminAuthModel) []base.BaseAuth {
	var data []base.BaseAuth
	for _, v := range auths {
		if v.Pid != pid {
			continue
		}
		data = append(data, base.BaseAuth{
			Id:    v.ID,
			Pid:   v.Pid,
			Name:  v.Name,
			Child: authTree(v.ID, auths),
		})
	}
	return data
}
