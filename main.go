package main
import (

	"net/http"

	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"fmt"
)

func hello(w http.ResponseWriter , r* http.Request)  {
	fmt.Fprintf(w, "hello")
}



func main() {
	//google 内置的多路器
	//muxold := http.NewServeMux()
	//fs := http.FileServer(http.Dir("."))
	//muxold.Handle("/public/" ,fs)


	gorillaRoute := mux.NewRouter()
	//gorillaRoute.HandleFunc("api/{user:[0-9]+}" ,hello)
	gorillaRoute.PathPrefix("/public/").Handler(http.StripPrefix("/public/" , http.FileServer(http.Dir("public/"))))


	n := negroni.Classic()
	//n.Use(negroni.Wrap(gorillaRoute))

	n.UseHandler(gorillaRoute)
	n.Run(":8080")
}
