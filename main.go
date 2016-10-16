package main
import (

	"net/http"
	"github.com/codegangsta/negroni"
	"fmt"
	"iaskServer/routers"
	"log"

	"iaskServer/common"
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
	n.UseHandler(router)

	server := &http.Server{
		Addr:common.AppConfig.Server,
		Handler:n ,
	}
	log.Println("Listening ...........")
	server.ListenAndServe()
}
