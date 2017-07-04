package routers

import (
	"github.com/gorilla/mux"
	"iaskServer/controllers/Admins"
	"github.com/codegangsta/negroni"
	"iaskServer/common"
	"net/http"

)

func setAdminsRouter(router *mux.Router) *mux.Router  {
	adminRouter := mux.NewRouter().StrictSlash(true)
	adminRouter.HandleFunc("/QAdmins" ,Admins.Index)
	adminRouter.HandleFunc("/QAdmins/index" , Admins.Index)


	router.PathPrefix("/QAdmin").Handler(negroni.New(
		negroni.HandlerFunc(common.AdminsAuthorize) ,
		negroni.Wrap(adminRouter),
	))



	//auth := router.PathPrefix("/auth").Subrouter()
	router.Handle("/auth/admlogin" , http.HandlerFunc(Admins.ShowSignupForm)).Methods("GET")
	router.Handle("/auth/admlogin" , http.HandlerFunc(Admins.SubmitSignupForm)).Methods("POST")



	return router
}