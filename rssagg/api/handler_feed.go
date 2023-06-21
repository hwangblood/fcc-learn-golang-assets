package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/models"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/reswrapper"
)

func (apiCfg *ApiConfig) HandleCreateFeed(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	reswrapper.ResponseWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandleGetFeeds(
	w http.ResponseWriter,
	r *http.Request,
) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	reswrapper.ResponseWithJSON(w, 200, models.DatabaseFeedsToFeeds(feeds))
}
