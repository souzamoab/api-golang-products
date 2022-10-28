package migrations

import (
	"github.com/souzamoab/api-golang-products/src/app/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Product{})
}
