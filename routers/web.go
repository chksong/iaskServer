package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	web "iaskServer/controllers/Web"

)

func SetWebRoutes(router *mux.Router) *mux.Router {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/" , http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/test" , web.Test)

	router.HandleFunc("/index" , web.Index)


	return router
}