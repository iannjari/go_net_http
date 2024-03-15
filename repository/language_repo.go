package repository

import (
	"fmt"
	"go_net_http/database"
	"go_net_http/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LanguageRepository interface {
	Save(language *model.Language) (*model.Language, error)
	GetById(id *uuid.UUID) (error, *model.Language)
	Delete(id *uuid.UUID) error
	// TODO: have a query/pagination impl
	Query() error
}

type repo struct {
}

var db *gorm.DB

func NewLanguageRepository() LanguageRepository {
	db = database.Database
	return &repo{}
}

func (*repo) Save(language *model.Language) (*model.Language, error) {

	var err error

	tx := db.Begin()

	err = db.Create(&language).Error

	if err != nil {
		tx.Rollback()
		return language, fmt.Errorf("Could not create entity. Error: %w", err.Error)
	}
	tx.Commit()
	return language, nil
}

func (*repo) GetById(id *uuid.UUID) (error, *model.Language) {
	return nil, nil
}

func (*repo) Delete(id *uuid.UUID) error {
	return nil
}

// TODO: have a query/pagination impl
func (*repo) Query() error {
	return nil
}
