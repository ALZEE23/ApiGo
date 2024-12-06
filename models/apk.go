package models

import (
	"gorm.io/gorm"
)

type Apk struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:VARCHAR(30);not null;default:null"`
	Cover       string  `json:"cover"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Game        *string `json:"game"`
	Footage     string  `json:"footage"`
}
