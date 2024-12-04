package models

import (
	"gorm.io/gorm"
)

type Apk struct {
	gorm.Model
	Name        string  `json:"name" gorm:"unique"`
	Cover       string  `json:"cover"`
	Title       string  `json:"title" gorm:"unique"`
	Description string  `json:"description"`
	Game        *string `json:"game" gorm:"unique"`
	Footage     string  `json:"footage"`
}
