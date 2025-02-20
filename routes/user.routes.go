package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)

	r.Get("/", controllers.GetUser)
	r.Patch("/", controllers.UpdateUser)

	return r
}
