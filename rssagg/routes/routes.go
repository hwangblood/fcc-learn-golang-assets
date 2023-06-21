package routes

import "github.com/hwangblood/fcc-learn-golang-assets/rssagg/api"

type RoutesCaller struct {
	apiCfg *api.ApiConfig
}

func New(apiCfg *api.ApiConfig) RoutesCaller {
	return RoutesCaller{
		apiCfg: apiCfg,
	}
}
