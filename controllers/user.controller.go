package controllers

import (
	"encoding/json"
	"net/http"

	"auth/config"
	"auth/models"
)

// GetUser - Fetch a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}
	config.Respond(w, http.StatusOK, user)

}

// UpdateUser - Update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var profile struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user, ok := r.Context().Value("user").(models.User)

	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	user.Name = profile.Name
	user.Age = &profile.Age

	config.DB.Save(&user)
	user.Password = ""
	config.Respond(w, http.StatusOK, user)
}
