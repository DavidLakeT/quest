package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	ID        uint
	URL       string
	Title     string `gorm:"unique;index"`
	Metadata  string
	Validated bool
	CitizenID uint
}
