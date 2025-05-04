package seeder

import (
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunRoleSeeder(db *gorm.DB){
	roles := []models.Role{
		{ 
			Code: constants.ROLE_ADMIN,
			Name: "Administrator",
		},
		{ 
			Code: constants.ROLE_USER,
			Name: "User",
		},
	}

	for _, role := range roles {
		err := db.FirstOrCreate(&role,models.Role{Code: role.Code}).Error
		if err != nil {
			logrus.Errorf("failed to seed role: %v", err)
			panic(err)
		}
		logrus.Infof("role %s successfully seededd", role.Code)
	}
}