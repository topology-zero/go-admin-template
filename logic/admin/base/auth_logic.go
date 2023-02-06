package base

import (
	"admin_template/model"
	"admin_template/query"
	"admin_template/types/admin/base"
)

func Auth() (resp base.BaseAuthResponse, err error) {
	authModel := query.AuthModel
	auths, _ := authModel.Find()
	resp.Tree = authTree(0, auths)
	return
}

func authTree(pid int, auths []*model.AuthModel) []base.BaseAuth {
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
