package models

import "gorm.io/gorm"

type Gratitude struct {
	gorm.Model
	Name        string
	Description string
}
