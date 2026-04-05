package middleware

import (
	"net/http"
	"strings"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, "Missing header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(auth, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid format", http.StatusUnauthorized)
			return
		}

		if parts[1] != "secret123" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}