package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Site entity
	router = SetSiteRoutes(router)
	// Routes for the SmokeTest entity
	router = SetSmokeTestRoutes(router)
	// Routes for the TestUrl entity
	router = SetUrlRoutes(router)
	return router
}
