package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/controllers"
)

func SetSiteRoutes(router *mux.Router) *mux.Router {
	siteRouter := mux.NewRouter()
	siteRouter.HandleFunc("/sites", controllers.CreateSite).Methods("POST")
	siteRouter.HandleFunc("/sites/{id}", controllers.UpdateSite).Methods("PUT")
	siteRouter.HandleFunc("/sites", controllers.GetSites).Methods("GET")
	siteRouter.HandleFunc("/sites/{id}", controllers.GetSiteById).Methods("GET")
	siteRouter.HandleFunc("/sites/users/{id}", controllers.GetSitesByUser).Methods("GET")
	siteRouter.HandleFunc("/sites/{id}", controllers.DeleteSite).Methods("DELETE")
	router.PathPrefix("/sites").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(siteRouter),
	))
	return router
}
