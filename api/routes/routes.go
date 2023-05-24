package Routes

import (
	"categorymanagement/api/handlers"
	"categorymanagement/cache"
	"categorymanagement/middlewares"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {

	//.With(middlewares.Authuser)

	r.With(middlewares.DbContext).With(middlewares.Authuser).Route("/categories", func(r chi.Router) {
		r.Post("/add", handlers.AddCategory)
		r.Put("/update", handlers.UpdateCategory)
		r.Delete("/delete", handlers.DeleteCategory)
		r.With(cache.CheckCache).Get("/get", handlers.GetCategories)
	})
}
