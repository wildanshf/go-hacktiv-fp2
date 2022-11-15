package middleware

import (
	"fmt"
	"net/http"
	"project-2/config"
	"strings"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("authorization")
		if bearerToken == "" {
			http.Error(w, "authorization invalid", http.StatusUnauthorized)
			fmt.Println(1)
			return
		}

		splitedToken := strings.Split(bearerToken, " ")
		if len(splitedToken) != 2 {
			http.Error(w, "authorization invalid", http.StatusUnauthorized)
			fmt.Println(2)
			return
		}

		if splitedToken[0] != "Bearer" {
			http.Error(w, "authorization invalid", http.StatusUnauthorized)
			fmt.Println(3)
			return
		}

		if !config.VerifyToken(splitedToken[1]) {
			http.Error(w, "authorization invalid", http.StatusUnauthorized)
			fmt.Println(4)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetClaim(r *http.Request) *config.MyClaim {
	bearerToken := r.Header.Get("authorization")
	if bearerToken == "" {
		return nil
	}

	splitedToken := strings.Split(bearerToken, " ")

	claim := config.GetClaim(splitedToken[1])
	return claim
}
