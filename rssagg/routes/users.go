package routes

import "net/http"

func (routes *RoutesCaller) CreateUser() http.HandlerFunc {
	return routes.apiCfg.HandleCreateUser
}

func (routes *RoutesCaller) GetUser() http.HandlerFunc {
	apiCfg := routes.apiCfg
	return apiCfg.MiddlewareAuth(apiCfg.HandleGetUser)
}
