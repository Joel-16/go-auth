package main

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/config"
	"auth/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := config.EnvValidation(); err != nil {
		errors, _ := json.MarshalIndent(err, "", " ")
		log.Fatal(errors)
	}
	config.ConnectDB()
	r := chi.NewRouter()
	r.Mount("/auth", routes.AuthRoutes())
	r.Mount("/users", routes.UserRoutes())
	r.Mount("/blogs", routes.BlogRoutes())

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	server := &http.Server{
		Handler: r,
		Addr:    ":" + config.PORT,
	}
	log.Printf("Server running on port %s", config.PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
