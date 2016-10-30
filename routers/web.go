package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	web "iaskServer/controllers/Web"

)

func SetWebRoutes(router *mux.Router) *mux.Router {
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/" , http.FileServer(http.Dir("assets/"))))
	router.HandleFunc("/test" , web.Test)

	router.HandleFunc("/index" , web.Index)


	return router
}