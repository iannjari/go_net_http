package api

import (
	"encoding/json"
	"fmt"
	"go_net_http/api/model"
	"go_net_http/api/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	*mux.Router

	languages []model.Language
}

var dbClient *gorm.DB
var languageService *service.LanguageService

func NewServer(db *gorm.DB) *Server {
	dbClient = db
	languageService = service.NewLanguageService(dbClient)
	fmt.Println("💡 Using Gorilla Mux router...")
	s := Server{
		Router:    mux.NewRouter(),
		languages: []model.Language{},
	}
	fmt.Println(" ")
	fmt.Println("Configured server using Gorilla Mux")

	fmt.Println("Configuring routes...")
	fmt.Println(" ")
	s.routes()
	fmt.Println("Routes configured.")
	fmt.Println("✅ Server up, listening on port: 8080")
	return &s
}

func (s *Server) routes() {
	s.HandleFunc("/languages", s.listLanguages()).Methods("GET")
	s.HandleFunc("/language/{id}", s.deleteLanguage()).Methods("DELETE")
	s.HandleFunc("/language", s.createLanguage()).Methods("POST")
}

func (s *Server) createLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var l model.Language
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		l, serviceErr := languageService.CreateLanguage(&l)

		w.Header().Set("Content-Type", "application/json")
		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}
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
		idStr := mux.Vars(r)["id"]
		fmt.Println(idStr)
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
