package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/myschool-petproj/class-ms/handlers"
)

func InitVersionedRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/classes", func(r chi.Router) {
			r.Get("/", handlers.GetClasses)
			r.Post("/", handlers.CreateClass)

			r.Get("/{id}", handlers.GetClass)
			r.Patch("/{id}", handlers.UpdateClass)
			r.Delete("/{id}", handlers.DeleteClass)
		})
		r.Route("/profiles", func(r chi.Router) {
			r.Get("/", handlers.GetProfiles)
			r.Get("/{id}", handlers.GetProfile)
		})
	})
}
