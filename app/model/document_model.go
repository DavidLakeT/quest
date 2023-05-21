package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	ID     uint
	URL    string
	UserID uint
}
