package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/alainglez/go-smoke/common"
	"github.com/alainglez/go-smoke/controllers"
)

func SetUrlRoutes(router *mux.Router) *mux.Router {
	urlRouter := mux.NewRouter()
	urlRouter.HandleFunc("/urls", controllers.CreateUrl).Methods("POST")
	urlRouter.HandleFunc("/urls/{id}", controllers.UpdateUrl).Methods("PUT")
	urlRouter.HandleFunc("/urls/{id}", controllers.GetUrlById).Methods("GET")
	urlRouter.HandleFunc("/urls", controllers.GetUrls).Methods("GET")
	urlRouter.HandleFunc("/urls/sites/{id}", controllers.GetUrlsBySite).Methods("GET")
	urlRouter.HandleFunc("/urls/{id}", controllers.DeleteUrl).Methods("DELETE")
	router.PathPrefix("/urls").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(urlRouter),
	))
	return router
}
