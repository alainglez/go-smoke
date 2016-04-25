package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)
/*
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")

	userJson := `{"firstname": "Alain", "lastname": "Gonzalez", "email": "alain.gonzalez@xyz.com"}`

	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
	}
}
func TestUniqueEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")

	userJson := `{"firstname": "Alain", "lastname": "Gonzalez", "email": "alain.gonzalez@xyz.com"}`

	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Error("Bad Request expected, got: %d", w.Code)
	}
}
*/
func TestRegisterClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users/register", Register).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users/register", server.URL)
	userJson := `{"firstname": "Ileana", "lastname": "Gonzalez", "email": "ileana.gonzalez@xyz.com"}`
	request, err := http.NewRequest("POST", usersUrl, strings.NewReader(userJson))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
/*
func TestLoginClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", Login).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
*/

