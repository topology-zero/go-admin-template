package generate

import (
	"go-admin-template/model"
	"go-admin-template/svc"
	"go-admin-template/types/admin/generate"
)

// List 表结构列表
func List(ctx *svc.ServiceContext) (resp generate.GenerateListResponse, err error) {
	model.DB().Raw("show tables").Scan(&resp.Tables)
	return
}
