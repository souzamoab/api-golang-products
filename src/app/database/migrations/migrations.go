package migrations

import (
	"github.com/fabriciojlm/api-golang-users/src/app/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
