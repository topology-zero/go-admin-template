package generate

import (
	"admin_template/model"
	"admin_template/svc"
	"admin_template/types/admin/generate"
)

// List 表结构列表
func List(ctx *svc.ServiceContext) (resp generate.GenerateListResponse, err error) {
	model.DB().Raw("show tables").Scan(&resp.Tables)
	return
}
