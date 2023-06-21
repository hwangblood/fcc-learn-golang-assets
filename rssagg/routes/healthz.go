package routes

import (
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/api"
)

func (routes *RoutesCaller) Healthz() http.HandlerFunc {
	return api.HandleReadiness
}
