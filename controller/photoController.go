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

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		photo model.Photo
		claim = middleware.GetClaim(r)
	)

	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.UserID = claim.ID

	result, err := service.CreatePhoto(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoUrl,
		"user_id":    result.UserID,
		"created_at": result.CreatedAt,
	})
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		claim  = middleware.GetClaim(r)
		result = []map[string]interface{}{}
	)

	photos, err := service.GetPhotos(claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range photos {
		result = append(result, map[string]interface{}{
			"id":         v.ID,
			"title":      v.Title,
			"caption":    v.Caption,
			"photo_url":  v.PhotoUrl,
			"user_id":    v.UserID,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
			"user": map[string]interface{}{
				"email":    v.User.Email,
				"username": v.User.Username,
			},
		})
	}

	json.NewEncoder(w).Encode(result)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		intID, _ = strconv.Atoi(id)
		claim    = middleware.GetClaim(r)
		photo    = model.Photo{
			ID: intID,
		}
	)

	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.UpdatePhoto(photo, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoUrl,
		"user_id":    result.UserID,
		"updated_at": result.UpdatedAt,
	})
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		intID, _ = strconv.Atoi(id)
		claim    = middleware.GetClaim(r)
	)

	_, err := service.DeletePhoto(intID, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Your photo has been successfully deleted",
	})
}
