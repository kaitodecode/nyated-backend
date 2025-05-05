package models

type Folder struct {
	BaseModel
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Description *string `json:"description" gorm:"type:varchar(255);"`
	UserID		string  `json:"user_id" gorm:"type:varchar(50);not null"`
	User		User	`json:"user" gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

