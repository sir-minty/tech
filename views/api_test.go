package views_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ApiError "github.com/sir-minty/tech/errors"
	"github.com/sir-minty/tech/views"
)

func TestLoginHandler(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sql mock error: %s", err)
	}
	defer db.Close()

	c := views.Context{DB: db}

	tests := map[string]struct {
		method     string
		url        string
		statusCode int
		username   string
		password   string
		response   string
		rows       sqlmock.Rows
	}{
		"200 ok": {
			"POST", "/login", http.StatusOK, "a", "b", "",
			sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "a", "b"),
		},
		"401 missing username": {
			"POST", "/login", http.StatusUnauthorized, "", "b", ApiError.MissingUsername.GoString(),
			nil,
		},
		"401 missing password": {
			"POST", "/login", http.StatusUnauthorized, "a", "", ApiError.MissingPassword.GoString(),
			nil,
		},
		"401 invalid username": {
			"POST", "/login", http.StatusUnauthorized, "aa", "b", ApiError.UsernameDoesNotExist.GoString(),
			nil,
		},
		"401 invalid password": {
			"POST", "/login", http.StatusUnauthorized, "a", "bb", ApiError.InvalidCredentials.GoString(),
			sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "a", "b"),
		},
	}

	for test, p := range tests {
		// Populate mock DB with user data only for tests that actually make it down to the DB call
		if p.rows != nil {
			mock.ExpectQuery("^SELECT (.+) FROM user WHERE (.+)$").WillReturnRows(p.rows)
		}

		// Form parameters
		f := url.Values{}
		if p.username != "" {
			f.Set("username", p.username)
		}
		if p.password != "" {
			f.Set("password", p.password)
		}

		r, _ := http.NewRequest(p.method, p.url, strings.NewReader(f.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		c.LoginHandler(w, r)

		if w.Code != p.statusCode {
			t.Errorf("%s: Expected status code %d got %d", test, p.statusCode, w.Code)
		}
		if w.Body.String() != p.response {
			t.Errorf("%s: Expected message %s got %s", test, p.response, string(w.Body.Bytes()))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("%s: %s", test, err)
		}
	}
}
