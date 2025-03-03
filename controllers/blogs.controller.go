package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/config"
	"auth/models"

	"github.com/go-chi/chi/v5"
)

// GetUser - Fetch a single user by ID

func CreateBlogs(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	blog.UserID = user.ID
	if err := config.DB.Create(&blog).Error; err != nil {
		log.Println("Error creating blog:", err)
		config.Respond(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	config.Respond(w, http.StatusOK, blog)

}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid user")
		return
	}
	config.Respond(w, http.StatusOK, blogs)

}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	id := chi.URLParam(r, "id")
	user, ok := r.Context().Value("user").(models.User)

	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}
	if err := config.DB.Preload("Comments").Where(map[string]interface{}{"id": id, "user_id": user.ID}).First(&blog).Error; err != nil {
		log.Println("Error retrieving blog: ", err)
		config.Respond(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	config.Respond(w, http.StatusOK, blog)

}

// UpdateUser - Update a user
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	var payload struct {
		Title string `json:"name"`
		Body  string `json:"age"`
	}

	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user, ok := r.Context().Value("user").(models.User)

	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	if err := config.DB.Where(map[string]interface{}{"ID": id, "UserID": user.ID}).First(&blog).Error; err != nil {
		config.Respond(w, http.StatusUnauthorized, struct{}{})
		return
	}

	blog.Body = payload.Body
	blog.Title = payload.Title
	config.DB.Save(&blog)
	config.Respond(w, http.StatusOK, blog)
}
