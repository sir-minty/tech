package views

import (
	"database/sql"
	"encoding/json"
	"net/http"

	apiError "github.com/sir-minty/tech/errors"
	"github.com/sir-minty/tech/models"
)

type Context struct {
	DB *sql.DB
}

func NewContext(db *sql.DB) *Context {
	return &Context{db}
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
	resp := struct {
		Message string
	}{":D WHOOO!"}
	c.ReturnResponse(w, resp)
}

func (c *Context) ReturnError(w http.ResponseWriter, e apiError.ApiError) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(e)
}

func (c *Context) ReturnResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
