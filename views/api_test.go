package views_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sir-wiggles/auth/views"
)

func TestLoginHandler(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sql mock error: %s", err)
	}
	defer db.Close()

	api := views.Context{DB: db}

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).
		AddRow(1, "jeff", "asdf")
	mock.ExpectQuery("^SELECT (.+) FROM user WHERE (.+)$").WillReturnRows(rows)

	tests := map[string]struct {
		method     string
		url        string
		statusCode int
		username   string
		password   string
	}{
		"200 ok": {"POST", "/login", 200, "jeff", "asdf"},
		//		"401 username invalid": {"POST", "/login", 401, "x", "asdf"},
		//		"401 password invalid": {"POST", "/login", 401, "jeff", "x"},
		//		"401 password missing": {"POST", "/login", 401, "jeff", ""},
		//		"401 username missing": {"POST", "/login", 401, "", "x"},
		//		"404 wrong method":     {"GET", "/login", 404, "jeff", "asdf"},
	}

	for test, p := range tests {
		h := api.LoginHandler
		w := httptest.NewRecorder()

		f := url.Values{}
		f.Set("username", p.username)
		f.Set("password", p.password)

		r, _ := http.NewRequest(p.method, p.url, bytes.NewBuffer([]byte(f.Encode())))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		h(w, r)

		if w.Code != p.statusCode {
			t.Errorf("%s: Expected %d got %d", test, p.statusCode, w.Code)
		}

	}

}
