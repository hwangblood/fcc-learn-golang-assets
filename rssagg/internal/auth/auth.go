package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetAPIKey extracts the API key from the headers of the request

Example:

Authorization: APIKey {insert apikey here}
*/
func GetAPIkey(header http.Header) (string, error) {
	authVal := header.Get("Authorization")
	if authVal == "" {
		return "", errors.New("no authorization header found")
	}

	authVals := strings.Split(authVal, " ")
	if len(authVals) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if authVals[0] != "APIKey" {
		return "", errors.New("malformed first part of authorization header")
	}

	return authVals[1], nil
}
