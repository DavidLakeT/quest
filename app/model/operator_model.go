package model

import "gorm.io/gorm"

type Operator struct {
	gorm.Model

	ID       uint
	Name     string
	URL      string
	Citizens []Citizen
}
