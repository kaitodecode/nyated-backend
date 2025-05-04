package seeder

import (
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	dummyPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	var admin, user models.Role
	if err := db.Where("code = ?", constants.ROLE_ADMIN).First(&admin).Error; err != nil {
		logrus.Errorf("%s", err.Error())
	}

	if err := db.Where("code = ?", constants.ROLE_USER).First(&user).Error; err != nil {
		logrus.Errorf("%s", err.Error())
	}

	users := []models.User{
		{
			Name: "admin",
			Email: "admin@gmail.com",
			Password: string(dummyPassword),
			RoleID: admin.ID,
		},
		{
			Name: "customer",
			Email: "customer@gmail.com",
			Password: string(dummyPassword),
			RoleID: user.ID,
		},
	}

	for _, user := range users {
		err := db.FirstOrCreate(&user, models.User{Email: user.Email}).Error
		if err != nil {
			logrus.Errorf("Failed to seed data: %v", err)
		}
		logrus.Infof("User %s successfully seeded", user.Email)
	}
}