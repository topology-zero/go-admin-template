package internal

import "gorm.io/gen"

type Role interface {

	// GetRoleByAuthId 根据 auth_id 获取角色
	//
	// where(json_contains( auth, json_array( @authId )) )
	GetRoleByAuthId(authId int) *gen.T
}
