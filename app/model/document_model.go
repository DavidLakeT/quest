package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	URL       string
	Title     string `gorm:"unique;index"`
	Validated bool
	CitizenID uint
}
