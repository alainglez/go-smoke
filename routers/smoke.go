package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/controllers"
)

func SetSmokeTestRoutes(router *mux.Router) *mux.Router {
	smoketestRouter := mux.NewRouter()
	smoketestRouter.HandleFunc("/smoketests", controllers.RunSmokeTest).Methods("POST")
	//PUT not needed as RunSmokeTest updates the PassFail and URLResults and returns them as JSON
	//smoketest.HandleFunc("/smoketests/{id}", controllers.UpdateSmoketest).Methods("PUT")
	smoketestRouter.HandleFunc("/smoketests", controllers.GetSmokeTests).Methods("GET")
	smoketestRouter.HandleFunc("/smoketests/{id}", controllers.GetSmokeTestById).Methods("GET")
	smoketestRouter.HandleFunc("/smoketests/users/{id}", controllers.GetSmokeTestsByUser).Methods("GET")
	smoketest.HandleFunc("/smoketests/{id}", controllers.DeleteSmoketest).Methods("DELETE")
	router.PathPrefix("/smoketests").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(smoketestRouter),
	))
	return router
}
