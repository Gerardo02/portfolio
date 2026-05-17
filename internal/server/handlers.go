package server

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"

	"portafolio/cmd/web"
	"portafolio/internal/data"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Handler(web.HomeContent()).ServeHTTP(w, r)
		return
	}
	templ.Handler(web.Home()).ServeHTTP(w, r)
}

func (s *Server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Handler(web.AboutContent()).ServeHTTP(w, r)
		return
	}
	templ.Handler(web.About()).ServeHTTP(w, r)
}

func (s *Server) projectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Handler(web.ProjectsContent()).ServeHTTP(w, r)
		return
	}
	templ.Handler(web.Projects()).ServeHTTP(w, r)
}

func (s *Server) todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Handler(web.TodoContent()).ServeHTTP(w, r)
		return
	}
	templ.Handler(web.TodoApp()).ServeHTTP(w, r)
}

func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")

	if task == "" {
		http.Error(w, "no task found in request", http.StatusBadRequest)
		return
	}

	key := rand.Text()

	data.Tasks[key] = task

	web.NoTask().Render(r.Context(), w)

	templ.Handler(web.Task(key, task)).ServeHTTP(w, r)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	delete(data.Tasks, key)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted correctly blud"))
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Message string `json:"message"`
	}{Message: "Hello World"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
