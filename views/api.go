package views

import (
	"database/sql"
	"encoding/json"
	"net/http"

	apiError "github.com/sir-wiggles/auth/errors"
	"github.com/sir-wiggles/auth/models"
)

type Context struct {
	DB *sql.DB
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username == "" {
		c.ReturnError(w, apiError.MissingUsername)
		return
	} else if password == "" {
		c.ReturnError(w, apiError.MissingPassword)
		return
	}

	user := &models.User{}

	row := c.DB.QueryRow("SELECT * FROM user WHERE username=?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		c.ReturnError(w, apiError.UsernameDoesNotExist)
		return
	}

	if user.Password != password {
		c.ReturnError(w, apiError.InvalidCredentials)
		return
	}
}

func (c *Context) ReturnError(w http.ResponseWriter, e apiError.ApiError) {
	json.NewEncoder(w).Encode(e.Error)
	w.WriteHeader(e.StatusCode)
	return
}
