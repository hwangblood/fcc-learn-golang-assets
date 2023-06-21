package routes

import "net/http"

func (routes *RoutesCaller) CreateFeedFollow() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleCreateFeedFollow)
}
func (routes *RoutesCaller) GetFeedFollows() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleGetFeedFollows)
}
func (routes *RoutesCaller) DeleteFeedFollow() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleDeleteFeedFollow)
}
