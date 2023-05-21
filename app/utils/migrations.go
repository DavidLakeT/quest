package utils

import (
	model "quest/model"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&model.Citizen{})
	db.AutoMigrate(&model.Document{})
}
