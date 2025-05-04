package models

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	RoleID 	 string `json:"roleID" gorm:"type:varchar(40);not null"`
	Role	 Role	`json:"role" gorm:"foreignKey:RoleID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,"`
}
