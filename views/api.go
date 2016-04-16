package api

import (
	"encoding/json"
	"net/http"
)

var users = map[string]string{
	"jeff": "asdf",
}

type apiErr struct {
	Error string `json:"error"`
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	if username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(apiErr{Error: "No username provided."})
		return
	}

	password := r.PostFormValue("password")
	if password == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(apiErr{Error: "No password provided."})
		return
	}

	userPassword, ok := users[username]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(apiErr{Error: "Username does not exist."})
		return
	}

	if userPassword != password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(apiErr{Error: "Invalid user credentials."})
		return
	}
})
