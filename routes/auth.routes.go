package routes

import (
	"auth/controllers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/signup", controllers.Signup)
	r.Post("/signin", controllers.Signin)
	r.Post("/forgot-password", controllers.GetUser)

	return r
}
