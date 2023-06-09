package main

import (
	"fmt"
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/auth"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) meddlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIkey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Authentication failed: %v\n", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v\n", err))
			return
		}

		handler(w, r, user)
	}
}
