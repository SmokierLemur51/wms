package routes

import (
	"net/http"
    "fmt"

    "github.com/go-chi/jwtauth/v5"
)

func AuthTestHandler(w http.ResponseWriter, r *http.Request) error {
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "Authentication error:"+err.Error(), http.StatusUnauthorized)
		return nil
	}

	if token == nil {
		http.Error(w, "No valid token found", http.StatusUnauthorized)
		return nil
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user id claim", http.StatusUnauthorized)
		return nil
	}

	w.Write([]byte(fmt.Sprintf("protected area, welcome %s", userID)))

	return nil
}