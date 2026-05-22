package server

import (
	"net/http"

	"portafolio/cmd/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Handle("/assets/*", http.FileServer(http.FS(web.Files)))

	r.Group(func(r chi.Router) {
		r.Get("/", s.homeHandler)
		r.Get("/about", s.aboutHandler)
		r.Get("/todo", s.todoHandler)
		r.Get("/projects", s.projectsHandler)

		r.Post("/task", s.newTaskHandler)
		r.Delete("/task/{key}", s.deleteTask)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/login/{provider}", s.loginHandler)
			r.Get("/logout/{provider}", s.logoutHandler)
			r.Get("/{provider}/callback", s.callbackHandler)
		})

		r.Post("/task", s.newTaskHandler)
	})

	return r
}
