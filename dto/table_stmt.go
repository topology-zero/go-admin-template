package dto

type CreateTable struct {
	Table       string
	CreateTable string `gorm:"column:Create Table"`
}

type Cols struct {
	Name    string
	Type    string
	Comment string
	Default string
}
