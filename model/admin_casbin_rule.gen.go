// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAdminCasbinRuleModel = "admin_casbin_rule"

// AdminCasbinRuleModel mapped from table <admin_casbin_rule>
type AdminCasbinRuleModel struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	Ptype string `gorm:"column:ptype"`
	V0    string `gorm:"column:v0"`
	V1    string `gorm:"column:v1"`
	V2    string `gorm:"column:v2"`
	V3    string `gorm:"column:v3"`
	V4    string `gorm:"column:v4"`
	V5    string `gorm:"column:v5"`
	V6    string `gorm:"column:v6"`
	V7    string `gorm:"column:v7"`
}

// TableName AdminCasbinRuleModel's table name
func (*AdminCasbinRuleModel) TableName() string {
	return TableNameAdminCasbinRuleModel
}
