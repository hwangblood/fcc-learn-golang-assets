package routes

import (
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/api"
)

func (routes *RoutesCaller) Err() http.HandlerFunc {
	return api.HandleErr
}
