package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/models"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/reswrapper"
)

func (apiCfg *ApiConfig) HandleCreateFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	reswrapper.ResponseWithJSON(
		w, 201,
		models.DatabaseFeedFollowToFeedFollow(feedFollow),
	)
}

func (apiCfg *ApiConfig) HandleGetFeedFollows(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	reswrapper.ResponseWithJSON(
		w, 200,
		models.DatabaseFeedFollowsToFeedFollows(feedFollows),
	)
}

func (apiCfg *ApiConfig) HandleDeleteFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")

	feedFollowID, paramErr := uuid.Parse(feedFollowIDStr)
	if paramErr != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", paramErr))
		return
	}

	_, existErr := apiCfg.DB.GetFeedFollow(r.Context(), feedFollowID)
	if existErr != nil {
		reswrapper.ResponseWithError(w, 404, fmt.Sprintf("Couldn't find feed follow: %v", existErr))
		return
	}

	deleteErr := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if deleteErr != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", deleteErr))
		return
	}

	reswrapper.ResponseWithJSON(w, 200, struct{}{})
}
