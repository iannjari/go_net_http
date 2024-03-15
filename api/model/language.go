package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Id       uuid.UUID `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Speakers int       `json:"speakers"`
}
