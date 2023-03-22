// Code generated by goctl. DO NOT EDIT.
package user

type PathId struct {
	Id int `uri:"id"`
}

type UserListRequest struct {
	Page     int `form:"page" label:"分页"`       // 分页
	PageSize int `form:"pageSize" label:"每页条数"` // 每页条数
}

type UserListResponse struct {
	Total int        `json:"total"` // 总条数
	Data  []UserList `json:"data"`  // 具体数据
}

type UserList struct {
	Id       int    `json:"id"`
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Rolename string `json:"rolename"` // 角色名
	Status   int    `json:"status"`   // 状态 0:未启用 1:正常
	Phone    string `json:"phone"`    // 手机号
}

type UserDetailResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
	RoleId   int    `json:"roleId"`   // 角色ID
	Status   int    `json:"status"`   // 状态 0:未启用 1:正常
}

type UserAddRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`              // 用户名
	Realname string `json:"realname" binding:"required" label:"真实姓名"`             // 真实姓名
	Phone    string `json:"phone" binding:"required" label:"手机号"`                 // 手机号
	RoleId   int    `json:"roleId" binding:"required" label:"角色ID"`               // 角色ID
	Password string `json:"password" binding:"required,min=6,max=255" label:"密码"` // 密码
	Status   int    `json:"status" binding:"oneof=0 1" label:"状态 0:未启用 1:正常"`     // 状态 0:未启用 1:正常
}

type UserEditRequest struct {
	PathId
	Username string `json:"username" binding:"required" label:"用户名"`          // 用户名
	Realname string `json:"realname" binding:"required" label:"真实姓名"`         // 真实姓名
	Phone    string `json:"phone" binding:"required" label:"手机号"`             // 手机号
	RoleId   int    `json:"roleId" binding:"required" label:"角色ID"`           // 角色ID
	Password string `json:"password" label:"密码"`                              // 密码
	Status   int    `json:"status" binding:"oneof=0 1" label:"状态 0:未启用 1:正常"` // 状态 0:未启用 1:正常
}
