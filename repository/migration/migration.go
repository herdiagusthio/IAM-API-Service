package migration

import (
	"iam-api-service/repository/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &user.Role{})
}
