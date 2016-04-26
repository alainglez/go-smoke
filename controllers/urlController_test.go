package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateUrl(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/urls", CreateUrl).Methods("POST")
	urlJson := `{"urldns": "www.carnival.com", "description": "Search & book carnival cruises."}`
	req, err := http.NewRequest(
		"POST",
		"/urls",
		strings.NewReader(urlJson),
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
func TestGetUrls(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/urls", GetUrls).Methods("GET")
	req, err := http.NewRequest("GET", "/urls", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}
func TestGetUrlByIdClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/urls/{id}", GetUrlById).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	urlsUrl := fmt.Sprintf("%s/urls/{id}", server.URL)
	request, err := http.NewRequest("GET", urlsUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
func TestGetUrlBySiteClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/urls/sites/{id}", GetUrlsBySite).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	urlsUrl := fmt.Sprintf("%s/urls/sites/{id}", server.URL)
	request, err := http.NewRequest("GET", urlsUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
func TestUpdateUrlClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/urls/{id}", UpdateUrl).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()
	urlsUrl := fmt.Sprintf("%s/urls/{id}", server.URL)
	request, err := http.NewRequest("PUT", urlsUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
func TestDeleteUrlClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/urls/{id}", Register).Methods("DELETE")
	server := httptest.NewServer(r)
	defer server.Close()
	urlsUrl := fmt.Sprintf("%s/urls/{id}", server.URL)
	request, err := http.NewRequest("DELETE", urlsUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	// 204 No Content
	if res.StatusCode != 204 {
		t.Errorf("HTTP Status expected: 204, got: %d", res.StatusCode)
	}
}
