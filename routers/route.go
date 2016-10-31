package routers

import (
	"github.com/gorilla/mux"

)

func InitRouters() *mux.Router  {
	router  := mux.NewRouter().StrictSlash(true)

	router  = SetWebRoutes(router)
	router  = setAdminsRouter(router)

	return  router
}

