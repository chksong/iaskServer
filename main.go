package main
import (

	"net/http"
	"fmt"
	"iaskServer/routers"
	"log"

	"iaskServer/common"
	"github.com/gorilla/csrf"
	"github.com/codegangsta/negroni"
)

func hello(w http.ResponseWriter , r* http.Request)  {
	fmt.Fprintf(w, "hello")
}



func main() {
	//calls startup logic
	common.StartUp()

	//Get the mux route object
	router := routers.InitRouters()

	//add init web
	//Web.Init()

	n := negroni.Classic()

	n.Use(negroni.NewStatic(http.Dir("/tmp")))
	n.UseHandler(router)

	log.Println("Listening ........")
/*	server := &http.Server{
		Addr:common.AppConfig.Server,
		Handler:n ,
	}*/


	crsfKey := []byte("keep-it-secret-keep-it-safe-----")
	CSRF := csrf.Protect(crsfKey)

	//添加csrf 保护
	http.ListenAndServe(common.AppConfig.Server, CSRF(n))
}
