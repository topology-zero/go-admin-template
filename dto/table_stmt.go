package dto

type CreateTable struct {
	Table       string
	CreateTable string `gorm:"column:Create Table"`
}

type Cols struct {
	Name      string
	MysqlType string
	GoType    string
	Comment   string
	Default   string
}
