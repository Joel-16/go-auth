package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/config"
	"auth/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetUser - Fetch a single user by ID

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	id, _ := uuid.Parse(chi.URLParam(r, "id"))

	var payload struct {
		Body string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	comment.UserID = user.ID
	comment.BlogID = id
	comment.Body = payload.Body
	if err := config.DB.Create(&comment).Error; err != nil {
		log.Println("Error creating comment:", err)
		config.Respond(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	config.Respond(w, http.StatusOK, comment)

}

// func Getcomments(w http.ResponseWriter, r *http.Request) {
// 	var comments []models.comment

// 	if err := config.DB.Find(&comments).Error; err != nil {
// 		config.Respond(w, http.StatusBadRequest, "Invalid user")
// 		return
// 	}
// 	config.Respond(w, http.StatusOK, comments)

// }

// func Getcomment(w http.ResponseWriter, r *http.Request) {
// 	var comment models.comment
// 	id := chi.URLParam(r, "id")
// 	userID, ok := r.Context().Value("user_id").(uint)

// 	if !ok || userID == 0 {
// 		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
// 		return
// 	}
// 	if err := config.DB.Preload("Comments").Where(map[string]interface{}{"ID": id, "UserID": userID}).First(&comment, "id=?", id).Error; err != nil {
// 		log.Println("Error retrieving comment: ", err)
// 		config.Respond(w, http.StatusInternalServerError, "Internal server error")
// 		return
// 	}
// 	config.Respond(w, http.StatusOK, comment)

// }

// UpdateUser - Update a user
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	var payload struct {
		Body string `json:"age"`
	}

	blogId := chi.URLParam(r, "id")
	commentId := chi.URLParam(r, "comment_id")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		config.Respond(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user, ok := r.Context().Value("user_id").(models.User)

	if !ok {
		config.Respond(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	if err := config.DB.Where(map[string]interface{}{"ID": commentId, "UserID": user.ID, "BlogID": blogId}).First(&comment).Error; err != nil {
		config.Respond(w, http.StatusUnauthorized, struct{}{})
		return
	}

	comment.Body = payload.Body
	config.DB.Save(&comment)
	config.Respond(w, http.StatusOK, comment)
}
