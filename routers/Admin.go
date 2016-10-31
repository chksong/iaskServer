package routers

import (
	"github.com/gorilla/mux"
	"iaskServer/controllers/Admins"
	"github.com/codegangsta/negroni"
	"iaskServer/common"
)

func setAdminsRouter(router *mux.Router) *mux.Router  {
	adminRouter := mux.NewRouter().StrictSlash(true)
	adminRouter.HandleFunc("/QAdmins" ,Admins.Index)
	adminRouter.HandleFunc("/QAdmins/index" , Admins.Index)


	router.PathPrefix("/QAdmin").Handler(negroni.New(
		negroni.HandlerFunc(common.AdminsAuthorize) ,
		negroni.Wrap(adminRouter),
	))

	router.HandleFunc("/admlogin" , Admins.ShowSignupForm).Methods("GET")
	router.HandleFunc("/admlogin/post" , Admins.SubmitSignupForm).Methods("POST")

	return router
}