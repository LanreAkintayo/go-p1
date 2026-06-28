package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization header should look like this;
// Authorization: ApiKey {apikey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Authorization header is missing")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("Authorization header is not in the expected format")
	}
	return vals[1], nil

}
