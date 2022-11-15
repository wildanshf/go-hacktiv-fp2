package controller

import (
	"encoding/json"
	"net/http"
	"project-2/config"
	"project-2/middleware"
	"project-2/model"
	"project-2/service"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       result.ID,
		"age":      result.Age,
		"email":    result.Email,
		"username": result.Username,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginParam model.LoginParam

	err := json.NewDecoder(r.Body).Decode(&loginParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isAllowed, user := service.IsLoginAllowed(loginParam); isAllowed {
		token := config.CreateToken(user.ID)

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "email and password not match",
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		claim = middleware.GetClaim(r)
		user  = model.User{
			ID: claim.ID,
		}
	)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         result.ID,
		"email":      result.Email,
		"username":   result.Username,
		"age":        result.Age,
		"updated_at": result.UpdatedAt,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		claim = middleware.GetClaim(r)
	)

	_, err := service.DeleteUser(claim.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Your account has been successfully deleted",
	})
}
