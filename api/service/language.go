package service

import (
	"fmt"
	"go_net_http/api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LanguageService struct {
	db *gorm.DB
}

func NewLanguageService(database *gorm.DB) *LanguageService {
	return &LanguageService{
		database,
	}
}

func (s *LanguageService) CreateLanguage(language *model.Language) (model.Language, error) {

	var err error
	language.Id = uuid.New()

	tx := s.db.Begin()

	err = s.db.Create(&language).Error

	if err != nil {
		tx.Rollback()
		return *language, fmt.Errorf("Could not create entity. Error: %w", err.Error)
	}
	tx.Commit()
	return *language, nil
}
