package routers

import (
	"github.com/gorilla/mux"

	"iaskServer/common"
)

func InitRouters() *mux.Router  {
	router  := mux.NewRouter().StrictSlash(true)

	router  = SetWebRoutes(router)
	router  = setAdminsRouter(router)

	/**
	  *  修改服务器权限， 检查CSRF的安全性
	 */
	router.Methods("POST", "PUT" ,"DELETE").Handler(common.CSRF(router))


	return  router
}

