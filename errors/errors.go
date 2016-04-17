package errors

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ApiError struct {
	ErrorMsg   string `json:"error"`
	StatusCode int    `json:"-"`
}

func (a *ApiError) GoString() string {
	w := bytes.NewBufferString("")
	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Println(err)
	}
	return w.String()
}

func (a *ApiError) Error() string {
	return a.GoString()
}

var (
	MissingUsername = ApiError{
		ErrorMsg:   "No username provided.",
		StatusCode: http.StatusUnauthorized,
	}

	MissingPassword = ApiError{
		ErrorMsg:   "No password provided.",
		StatusCode: http.StatusUnauthorized,
	}

	UsernameDoesNotExist = ApiError{
		ErrorMsg:   "Username does not exist.",
		StatusCode: http.StatusUnauthorized,
	}

	InvalidCredentials = ApiError{
		ErrorMsg:   "Invalid user credentials.",
		StatusCode: http.StatusUnauthorized,
	}
)
