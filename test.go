package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
)

func main()  {
	//golang 内置的服务器 单独使用没有问题
	//muxold := http.NewServeMux()
	//fs := http.FileServer(http.Dir("."))
	//muxold.Handle("/public/" ,fs)


	gorillaRoute := mux.NewRouter().StrictSlash(true)
	//gorillaRoute.HandleFunc("api/{user:[0-9]+}" ,hello)
	gorillaRoute.PathPrefix("/static/").Handler(http.StripPrefix("/static/" , http.FileServer(http.Dir("static/"))))


	n := negroni.Classic()
	//n.Use(negroni.Wrap(gorillaRoute))

	n.UseHandler(gorillaRoute)
	n.Run(":8080")
}
