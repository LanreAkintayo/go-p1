package main

import (
	"fmt"
	"net/http"

	"github.com/LanreAkintayo/go-p1/internal/auth"
	"github.com/LanreAkintayo/go-p1/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error getting API key: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error getting user by API key: %v", err))
			return
		}
		handler(w, r, user)
	}
}
