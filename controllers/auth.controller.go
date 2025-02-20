package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/config"
	"auth/models"
)

// CreateUser - Create a new user
func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, "email = ?", user.Email)
	if result.RowsAffected > 0 {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	user.Password = ""

	config.Respond(w, http.StatusCreated, user)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		config.Respond(w, http.StatusBadRequest, "Inavlid request Body")
		return
	}

	config.DB.First(&user, "email=?", credentials.Email)
	if user.ID == 0 {
		config.Respond(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	valid := user.CheckPasswordHash(credentials.Password, user.Password)

	if !valid {
		config.Respond(w, http.StatusBadRequest, "Invalid credentials")
		return
	}
	user.Password = ""
	token, err := user.GenerateToken(user.ID, config.JWT_SECRET)
	if err != nil {
		config.Respond(w, http.StatusInternalServerError, "Internal Server Error")
		log.Printf("Failed to generate jwt token")
	}

	config.Respond(w, http.StatusOK,
		struct {
			User  models.User `json:"user"`
			Token string      `json:"token"`
		}{
			User: user, Token: token,
		})
}

// // GetUser - Fetch a single user by ID
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	var user models.User

// 	if err := config.DB.First(&user, id).Error; err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// // UpdateUser - Update a user
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	var user models.User

// 	if err := config.DB.First(&user, id).Error; err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	json.NewDecoder(r.Body).Decode(&user)
// 	config.DB.Save(&user)
// 	json.NewEncoder(w).Encode(user)
// }

// // DeleteUser - Delete a user
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	config.DB.Delete(&models.User{}, id)
// 	w.WriteHeader(http.StatusNoContent)
// }
