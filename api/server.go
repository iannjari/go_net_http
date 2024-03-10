package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Language struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Speakers int       `json:"speakers"`
}

type Server struct {
	*mux.Router

	languages []Language
}

func NewServer() *Server {
	s := Server{
		Router:    mux.NewRouter(),
		languages: []Language{},
	}

	return &s
}

func (s *Server) createLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var l Language
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		l.Id = uuid.New()
		s.languages = append(s.languages, l)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(l); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Server) listLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.languages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Server) deleteLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, language := range s.languages {
			if id == language.Id {
				s.languages = append(s.languages[:i], s.languages[i+1:]...)
				break
			}
		}

	}
}
