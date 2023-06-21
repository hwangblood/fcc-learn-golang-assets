package routes

import "net/http"

func (routes *RoutesCaller) GetPostsForUser() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleGetPostsForUser)
}
