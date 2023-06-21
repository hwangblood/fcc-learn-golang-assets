package api

import (
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/reswrapper"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	reswrapper.ResponseWithJSON(w, 200, struct{}{})
}
