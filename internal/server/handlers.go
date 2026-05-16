package server

import (
	"net/http"

	"portafolio/cmd/web"

	"github.com/a-h/templ"
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
