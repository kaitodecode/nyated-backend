package models

type Note struct {
	BaseModel
	Title string `json:"title" gorm:"type:varchar(100);not null"`
	Content string `json:"content" gorm:"type:varchar(1000);not null"`
	FolderID string	`json:"folder_id" gorm:"type:varchar(50);not null"`
	Folder 	Folder `json:"folder" gorm:"foreignKey:FolderID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}