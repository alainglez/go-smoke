package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/data"
	"github.com/alainglez/go-smoke/models"
)

// Handler for HTTP Post - "/urls"
// Insert a new Url document for a SiteId
func CreateUrl(w http.ResponseWriter, r *http.Request) {
	var dataResource UrlResource
	// Decode the incoming Url json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Url data", 500)
		return
	}
	urlModel := dataResource.Data
	url := &models.SiteUrl{
		SiteId:      bson.ObjectIdHex(urlModel.SiteId),
		Description: urlModel.Description,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	//Insert a url document
	repo := &data.UrlRepository{c}
	repo.Create(url)
	if j, err := json.Marshal(url); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/urls/sites/{id}
// Returns all Urls documents under a SiteId
func GetUrlsBySite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	urls := repo.GetBySite(id)
	j, err := json.Marshal(UrlsResource{Data: urls})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/urls"
// Returns all Url documents
func GetUrls(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	urls := repo.GetAll()
	j, err := json.Marshal(UrlsResource{Data: urls})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/urls/{id}"
// Returns a single Url document by id
func GetUrlById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	url, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}
	if j, err := json.Marshal(url); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Handler for HTTP Put - "/urls/{id}"
// Update an existing Url document
func UpdateUrl(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource UrlResource
	// Decode the incoming Url json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Url data", 500)
		return
	}
	urlModel := dataResource.Data
	url := &models.SiteUrl{
		Id:          id,
		Description: urlModel.Description,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	//Update url document
	if err := repo.Update(url); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// Handler for HTTP Delete - "/urls/{id}"
// Delete an existing Url document
func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	//Delete a url document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
