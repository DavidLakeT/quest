package model

import "gorm.io/gorm"

type Citizen struct {
	gorm.Model

	ID           uint
	Name         string
	Address      string
	Email        string
	OperatorID   int
	OperatorName string
	Documents    []Document
}
