package api

import (
	"fmt"
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/reswrapper"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetAPIkey(r.Header)
		if err != nil {
			reswrapper.ResponseWithError(w, 403, fmt.Sprintf("Authentication failed: %v\n", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v\n", err))
			return
		}

		handler(w, r, user)
	}
}
