package service

import (
	"go_net_http/model"
	"go_net_http/repository"

	"github.com/google/uuid"
)

type LanguageService interface {
	Validate(language *model.Language) error
	Create(language *model.Language) (*model.Language, error)
	Fetch(id *uuid.UUID) (error, *model.Language)
	Delete(id *uuid.UUID) error
	// TODO: have a query/pagination impl
	QueryLanguages() error
}

type service struct {
}

var repo repository.LanguageRepository

func NewLanguageService() LanguageService {
	repo = repository.NewLanguageRepository()
	return &service{}
}

func (*service) Create(language *model.Language) (*model.Language, error) {
	return repo.Save(language)
}

func (*service) Validate(language *model.Language) error {
	return nil
}

func (*service) Fetch(id *uuid.UUID) (error, *model.Language) {
	return nil, nil
}

func (*service) Delete(id *uuid.UUID) error {
	return nil
}

// TODO: have a query/pagination impl
func (*service) QueryLanguages() error {
	return nil
}
