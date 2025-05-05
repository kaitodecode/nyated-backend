package seeder

import (
	"fmt"

	"github.com/kaitodecode/nyated-backend/domain/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunFolderSeeder(db *gorm.DB) {
	var user models.User
	err := db.Where("email = ?", "demo@gmail.com").First(&user).Error 
	
	if err != nil {
		logrus.Error(err)
	}

	desc := "Some description"
	var folders []models.Folder 
	for i := range 20 {
		folders = append(folders, models.Folder{
			Name: fmt.Sprintf("folder ke-%d",i+1),
			Description: &desc,
			UserID: user.ID,
		})
	}
	// err = db.Create(&folders).Error
	if err != nil {
		logrus.Error(err)
	}
}