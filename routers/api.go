package routers

import (
	"github.com/gorilla/mux"

	"iaskServer/iaskiAPI"
)

func SetApiRouter(router *mux.Router) *mux.Router {
 	router.HandleFunc("/api/json", iaskiAPI.TestJson)


	return router
}