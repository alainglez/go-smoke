package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/test/dockertest"
	"gopkg.in/mgo.v2"
)

func TestCreateSite(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/sites", CreateSite).Methods("POST")
	siteJson := `{"sitedns": "www.carnival.com", "description": "Search & book carnival cruises."}`
	req, err := http.NewRequest(
		"POST",
		"/sites",
		strings.NewReader(siteJson),
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
func TestGetSites(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/sites", GetSites).Methods("GET")
	req, err := http.NewRequest("GET", "/sites", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}
func TestGetSiteByIdClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/sites/{id}", GetSiteById).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	sitesUrl := fmt.Sprintf("%s/sites/{id}", server.URL)
	request, err := http.NewRequest("GET", sitesUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
func TestGetSiteByUserClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/sites/users/{id}", GetSitesByUser).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	sitesUrl := fmt.Sprintf("%s/sites/users/{id}", server.URL)
	request, err := http.NewRequest("GETT", sitesUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
func TestUpdateSiteClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/sites/{id}", UpdateUrl).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()
	sitesUrl := fmt.Sprintf("%s/sites/{id}", server.URL)
	request, err := http.NewRequest("PUT", sitesUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
func TestDeleteSiteClient(t *testing.T) {
	t.Parallel()
	r := mux.NewRouter()
	r.HandleFunc("/sites/{id}", Register).Methods("DELETE")
	server := httptest.NewServer(r)
	defer server.Close()
	sitesUrl := fmt.Sprintf("%s/sites/{id}", server.URL)
	request, err := http.NewRequest("DELETE", sitesUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	// 204 No Content
	if res.StatusCode != 204 {
		t.Errorf("HTTP Status expected: 204, got: %d", res.StatusCode)
	}
}
