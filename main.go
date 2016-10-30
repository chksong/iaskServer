package main
import (

	"net/http"
	"github.com/codegangsta/negroni"
	"fmt"
	"iaskServer/routers"
	"log"

	"iaskServer/common"
	"github.com/gorilla/csrf"
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

	//server := &http.Server{
	//	Addr:common.AppConfig.Server,
	//	Handler:n ,
	//}
	log.Println("Listening ...........")
	//server.ListenAndServe()

	http.ListenAndServe(common.AppConfig.Server,
		csrf.Protect([]byte("f943c6fe3b10ca918d17a327fc836c5a"))(n))
}
