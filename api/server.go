package api

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("💡 Using Gorilla Mux router...")
	s := Server{
		Router:    mux.NewRouter(),
		languages: []Language{},
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
		idStr := mux.Vars(r)["id"]
		fmt.Println(idStr)
		idU, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, language := range s.languages {
			if idU == language.Id {
				s.languages = append(s.languages[:i], s.languages[i+1:]...)
				break
			}
		}

	}
}
