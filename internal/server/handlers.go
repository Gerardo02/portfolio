package server

import (
	"crypto/rand"
	"net/http"

	"portafolio/cmd/web"
	"portafolio/internal/data"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

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
