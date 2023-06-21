package routes

import "net/http"

func (routes *RoutesCaller) CreateFeed() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleCreateFeed)
}

func (routes *RoutesCaller) GetFeeds() http.HandlerFunc {
	return routes.apiCfg.HandleGetFeeds
}
