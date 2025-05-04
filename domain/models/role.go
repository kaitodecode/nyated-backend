package models

type Role struct {
	BaseModel
	Code string `json:"code" gorm:"type:varchar(50);not null"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}