package errors

import "net/http"

type ApiError struct {
	Error      string `json:"error"`
	StatusCode int    `json:"-"`
}

var (
	MissingUsername = ApiError{
		Error:      "No username provided.",
		StatusCode: http.StatusUnauthorized,
	}

	MissingPassword = ApiError{
		Error:      "No password provided.",
		StatusCode: http.StatusUnauthorized,
	}

	UsernameDoesNotExist = ApiError{
		Error:      "Username does not exist.",
		StatusCode: http.StatusUnauthorized,
	}

	InvalidCredentials = ApiError{
		Error:      "Invalid user credentials.",
		StatusCode: http.StatusUnauthorized,
	}
)
