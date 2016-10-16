package routers

import (
	"github.com/gorilla/mux"
	"iaskServer/controllers/Admins"
	"github.com/codegangsta/negroni"
	"iaskServer/common"
)

func setAdminsRouter(router *mux.Router) *mux.Router  {
	adminRouter := mux.NewRouter()
	adminRouter.HandleFunc("/QAdmins" ,Admins.Index)
	adminRouter.HandleFunc("/QAdmins/index" , Admins.Index)
	adminRouter.HandleFunc("QAdmins/login" , Admins.Login)

	router.PathPrefix("/QAdmin").Handler(negroni.New(
		negroni.HandlerFunc(common.AdminsAuthorize) ,
		negroni.Wrap(adminRouter),
	))

	return router
}