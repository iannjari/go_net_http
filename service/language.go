package service

import (
	"errors"
	"go_net_http/model"
	"go_net_http/repository"

	"github.com/google/uuid"
)

type LanguageService interface {
	Validate(language *model.Language) error
	Create(language *model.Language) (*model.Language, error)
	Fetch(id *uuid.UUID) (*model.Language, error)
	Delete(id uuid.UUID) error
	// TODO: have a query/pagination impl
	QueryLanguages() (*[]model.Language, error)
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
	if len(language.Name) == 0 {
		return errors.New("the entity language does not have a name")
	}
	if language.Speakers == 0 {
		return errors.New("the entity language shpould have at least one speaker")
	}
	return nil
}

func (*service) Fetch(id *uuid.UUID) (l *model.Language, err error) {

	l, err = repo.GetById(id)
	if l == nil {
		err = errors.New("language with key %s not found" + id.String())
	}
	return
}

func (*service) Delete(id uuid.UUID) error {
	return repo.Delete(id)
}

// TODO: have a query/pagination impl
func (*service) QueryLanguages() (languages *[]model.Language, err error) {
	languages, err = repo.Query()
	return
}
