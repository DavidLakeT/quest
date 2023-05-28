package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	URL       string
	Title     string
	Validated bool
	CitizenID uint
}
