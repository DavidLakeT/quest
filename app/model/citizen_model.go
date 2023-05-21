package model

import "gorm.io/gorm"

type Citizen struct {
	gorm.Model

	ID        uint
	Documents []Document
}
