package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/controllers"
)

func SetSiteRoutes(router *mux.Router) *mux.Router {
	siteRouter := mux.NewRouter()
	siteRouter.HandleFunc("/sites", controllers.createSite).Methods("POST")
	siteRouter.HandleFunc("/sites/{id}", controllers.updateSite).Methods("PUT")
	siteRouter.HandleFunc("/sites", controllers.getSites).Methods("GET")
	siteRouter.HandleFunc("/sites/{id}", controllers.getSiteById).Methods("GET")
	siteRouter.HandleFunc("/sites/users/{id}", controllers.getSitesByUser).Methods("GET")
	siteRouter.HandleFunc("/sites/{id}", controllers.deleteSite).Methods("DELETE")
	router.PathPrefix("/sites").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(siteRouter),
	))
	return router
}
