package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/go-chi/chi/v5"
)

func BlogRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)

	r.Get("/", controllers.GetBlogs)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", controllers.CreateBlogs)
		r.Get("/{id}", controllers.GetBlog)
		r.Patch("/{id}", controllers.UpdateBlog)
		r.Post("/{id}/comments", controllers.CreateComment)
		r.Patch("/{id}/comments/{comment_id}", controllers.UpdateComment)
	})
	return r
}
