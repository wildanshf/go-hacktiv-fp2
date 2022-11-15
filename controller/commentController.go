package controller

import (
	"encoding/json"
	"net/http"
	"project-2/middleware"
	"project-2/model"
	"project-2/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AddComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		comment model.Comment
		claim   = middleware.GetClaim(r)
	)

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.UserID = claim.ID

	result, err := service.CreateComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         result.ID,
		"message":    result.Message,
		"photo_id":   result.PhotoID,
		"user_id":    result.UserID,
		"created_at": result.CreatedAt,
	})
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		claim  = middleware.GetClaim(r)
		result = []map[string]interface{}{}
	)

	comments, err := service.GetComments(claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range comments {
		result = append(result, map[string]interface{}{
			"id":         v.ID,
			"message":    v.Message,
			"photo_id":   v.PhotoID,
			"user_id":    v.UserID,
			"updated_at": v.UpdatedAt,
			"created_at": v.CreatedAt,
			"user": map[string]interface{}{
				"id":       v.User.ID,
				"email":    v.User.Email,
				"username": v.User.Username,
			},
			"photo": map[string]interface{}{
				"id":        v.Photo.ID,
				"title":     v.Photo.Title,
				"caption":   v.Photo.Caption,
				"photo_url": v.Photo.PhotoUrl,
				"user_id":   v.Photo.UserID,
			},
		})
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		intID, _ = strconv.Atoi(id)
		claim    = middleware.GetClaim(r)
		comment  = model.Comment{
			ID: intID,
		}
	)

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.UpdateComment(comment, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         result.ID,
		"message":    result.Message,
		"updated_at": result.UpdatedAt,
	})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		intID, _ = strconv.Atoi(id)
		claim    = middleware.GetClaim(r)
	)

	_, err := service.DeleteComment(intID, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Your comment has been successfully deleted",
	})
}
