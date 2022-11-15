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

func AddSocialMedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		socialMedia model.SocialMedia
		claim       = middleware.GetClaim(r)
	)

	err := json.NewDecoder(r.Body).Decode(&socialMedia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	socialMedia.UserID = claim.ID

	result, err := service.AddPSocialMedia(socialMedia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":               result.ID,
		"name":             result.Name,
		"social_media_url": result.SocialMediaUrl,
		"user_id":          result.UserID,
		"created_at":       result.CreatedAt,
	})
}

func GetSocialMedias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		claim  = middleware.GetClaim(r)
		result = []map[string]interface{}{}
	)

	socialMedias, err := service.GetSocialMedias(claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo, err := service.GetPhotoDetail(claim.ID)
	var photo_url string
	if err != nil {
		photo_url = "not found"
	} else {
		photo_url = photo.PhotoUrl
	}

	for _, v := range socialMedias {
		result = append(result, map[string]interface{}{
			"id":               v.ID,
			"name":             v.Name,
			"social_media_url": v.SocialMediaUrl,
			"created_at":       v.CreatedAt,
			"updated_at":       v.UpdatedAt,
			"user": map[string]interface{}{
				"id":                v.User.ID,
				"username":          v.User.Username,
				"profile_image_url": photo_url,
			},
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"social_medias": result,
	})
}

func UpdateSocialMedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars        = mux.Vars(r)
		id          = vars["id"]
		intID, _    = strconv.Atoi(id)
		claim       = middleware.GetClaim(r)
		socialMedia = model.SocialMedia{
			ID: intID,
		}
	)

	err := json.NewDecoder(r.Body).Decode(&socialMedia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.UpdateSocialMedia(socialMedia, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":               result.ID,
		"name":             result.Name,
		"social_media_url": result.SocialMediaUrl,
		"user_id":          result.UserID,
		"updated_at":       result.UpdatedAt,
	})
}

func DeleteSocialMedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		intID, _ = strconv.Atoi(id)
		claim    = middleware.GetClaim(r)
	)

	_, err := service.DeleteSocialMedia(intID, claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Your social media has been successfully deleted",
	})
}
