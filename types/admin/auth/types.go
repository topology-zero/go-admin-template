// Code generated by goctl. DO NOT EDIT.
package auth

type PathId struct {
	Id int `uri:"id"`
}

type AuthListResponse struct {
	Tree []Auth `json:"tree"` // 节点树
}

type Auth struct {
	Id     int    `json:"id"`
	Pid    int    `json:"pid"`
	Key    string `json:"key"`    // 权限标识
	Name   string `json:"name"`   // 节点名
	IsMenu int    `json:"isMenu"` // 是否是菜单栏 0：否 1：是
	Api    string `json:"api"`    // 接口
	Action string `json:"action"` // 操作方法
	Child  []Auth `json:"child"`  // 子节点
}

type AuthAddRequest struct {
	Pid    int    `json:"pid" binding:"min=0"`
	Name   string `json:"name" binding:"required" label:"节点名"`               // 节点名
	Key    string `json:"key" binding:"required" label:"权限标识"`               // 权限标识
	IsMenu int    `json:"isMenu" binding:"oneof=0 1" label:"是否是菜单栏 0：否 1：是"` // 是否是菜单栏 0：否 1：是
	Api    string `json:"api" label:"接口"`                                    // 接口
	Action string `json:"action" label:"操作方法"`                               // 操作方法
}

type AuthEditRequest struct {
	PathId
	AuthAddRequest
}
