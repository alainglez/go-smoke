package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/controllers"
)

func SetSmokeTestRoutes(router *mux.Router) *mux.Router {
	smoketestRouter := mux.NewRouter()
	smoketestRouter.HandleFunc("/smoketests", controllers.CreateSmokeTest).Methods("POST")
	smoketestRouter.HandleFunc("/smoketests", controllers.GetSmokeTests).Methods("GET")
	smoketestRouter.HandleFunc("/smoketests/{id}", controllers.GetSmokeTestById).Methods("GET")
	smoketestRouter.HandleFunc("/smoketests/users/{id}", controllers.GetSmokeTestsByUser).Methods("GET")
	router.PathPrefix("/smoketests").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(smoketestRouter),
	))
	return router
}
