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

func (apiCfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	reswrapper.ResponseWithJSON(w, 201, models.DatabaseUserToUser(user))
}

// this should be an authenticated endpoint
func (apiCfg *ApiConfig) HandleGetUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	reswrapper.ResponseWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetPostsForUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10, // TODO: limit from url query parameters
	})
	if err != nil {
		reswrapper.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}
	reswrapper.ResponseWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
