package util

import (
	"net/http"
)

func CheckHttpStatusCode(resp *http.Response) bool {
	switch resp.StatusCode {
	case http.StatusInternalServerError:
		return false
	case http.StatusBadGateway:
		return false
	case http.StatusBadRequest:
		return false
	}
	return true
}
