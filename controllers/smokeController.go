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

// Handler for HTTP Post - "/smoketests"
// Insert a new SmokeTest document 
// Runs the smoke test
// Update the SmokeTest document
// Returns status code results by url and overall pass fail in JSON response
func RunSmokeTest(w http.ResponseWriter, r *http.Request) {
	var dataResource SmokeTestResource
	// Decode the incoming SmokeTest json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid SmokeTest data",
			500,
		)
		return
	}
	smoketest := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	// Insert a smoketest document
	repo.Create(smoketest)
	// Get the urls for the siteid
	contexturls := NewContext()
	defer contexturls.Close()
	curls := context.DbCollection("urls")
	repourls := &data.UrlRepository{curls}
	testurls := repourls.GetBySite(smoketest.SiteId)
	// Run smoke test
	Smoke(smoketest,testurls)
	// Update the smoke test record with results
	if err := repo.Update(smoketest); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	// Marshal and return SmokeTest results in JSON response
	if j, err := json.Marshal(SmokeTestResource{Data: *smoketest}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOk)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/smoketests"
// Returns all SmokeTest documents
func GetSmokeTests(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	smoketests := repo.GetAll()
	j, err := json.Marshal(SmokeTestsResource{Data: smoketests})
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

// Handler for HTTP Get - "/smoketests/{id}"
// Returns a single SmokeTest document by id
func GetSmokeTestById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	smoketest, err := repo.GetById(id)
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
	if j, err := json.Marshal(smoketest); err != nil {
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

// Handler for HTTP Get - "/smoketests/users/{id}"
// Returns all SmokeTests created by a User
func GetSmokeTestsByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	smoketests := repo.GetByUser(user)
	j, err := json.Marshal(SmokeTestsResource{Data: smoketests})
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

// Handler for HTTP Put - "/smoketests/{id}"
// Update an existing SmokeTest document
// and Returns the PassFail and URLStatusCodes as JSON in http response
func UpdateSmokeTest(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource SmokeTestResource
	// Decode the incoming SmokeTest json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid SmokeTest data",
			500,
		)
		return
	}
	smoketest := &dataResource.Data
	smoketest.Id = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	// Update an existing SmokeTest document
	if err := repo.Update(smoketest); err != nil {
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

// Handler for HTTP Delete - "/smoketests/{id}"
// Delete an existing SmokeTest document
func DeleteSmokeTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("smoketests")
	repo := &data.SmokeTestRepository{c}
	// Delete an existing SmokeTest document
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
