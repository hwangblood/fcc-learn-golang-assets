package api

import (
	"net/http"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/reswrapper"
)

func HandleErr(w http.ResponseWriter, r *http.Request) {
	reswrapper.ResponseWithError(w, 400, "Something went wrong")
}
