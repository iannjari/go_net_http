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
	GetById(id *uuid.UUID) (*model.Language, error)
	Delete(id uuid.UUID) (err error)
	// TODO: have a query/pagination impl
	Query() (*[]model.Language, error)
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
		return language, fmt.Errorf("could not create entity. error: " + err.Error())
	}
	tx.Commit()
	return language, nil
}

func (*repo) GetById(id *uuid.UUID) (language *model.Language, error error) {
	db.First(&language, id.String)
	return
}

func (*repo) Delete(id uuid.UUID) (err error) {
	var lan model.Language
	lan.Id = id
	db.Delete(&lan)
	return
}

// TODO: have a query/pagination impl
func (*repo) Query() (languages *[]model.Language, err error) {
	db.Find(&languages)
	return
}
