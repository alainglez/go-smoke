package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/data"
)

// Handler for HTTP Post - "/sites"
// Insert a new Site document
func createSite(w http.ResponseWriter, r *http.Request) {
	var dataResource SiteResource
	// Decode the incoming Site json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Site data",
			500,
		)
		return
	}
	site := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	// Insert a site document
	repo.Create(site)
	if j, err := json.Marshal(SiteResource{Data: *site}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/sites"
// Returns all Site documents
func getSites(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	sites := repo.GetAll()
	j, err := json.Marshal(SitesResource{Data: sites})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// Handler for HTTP Get - "/sites/{id}"
// Returns a single Site document by id
func getSiteById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	site, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
	}
	if j, err := json.Marshal(site); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/sites/users/{id}"
// Returns all Sites created by a User
func getSitesByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	sites := repo.GetByUser(user)
	j, err := json.Marshal(SitesResource{Data: sites})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// Handler for HTTP Put - "/sites/{id}"
// Update an existing Site document
func updateSite(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource SiteResource
	// Decode the incoming Site json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Site data",
			500,
		)
		return
	}
	site := &dataResource.Data
	site.Id = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	// Update an existing Site document
	if err := repo.Update(site); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// Handler for HTTP Delete - "/sites/{id}"
// Delete an existing Site document
func deleteSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("sites")
	repo := &data.SiteRepository{c}
	// Delete an existing Site document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
