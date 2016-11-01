package main

import (
	//"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/csrf"
	"iaskServer/controllers/Admins"
)

func main444()  {
	////golang 内置的服务器 单独使用没有问题
	////muxold := http.NewServeMux()
	////fs := http.FileServer(http.Dir("."))
	////muxold.Handle("/public/" ,fs)
	//
	//
	//gorillaRoute := mux.NewRouter().StrictSlash(true)
	////gorillaRoute.HandleFunc("api/{user:[0-9]+}" ,hello)
	//gorillaRoute.PathPrefix("/static/").Handler(http.StripPrefix("/static/" , http.FileServer(http.Dir("static/"))))
	//
	//
	//n := negroni.Classic()
	////n.Use(negroni.Wrap(gorillaRoute))
	//
	//n.UseHandler(gorillaRoute)
	//n.Run(":8080")


}



func main() {
	CSRF := csrf.Protect(
	[]byte("a-32-byte-long-key-goes-here"),
	)

	r := mux.NewRouter()
	r.HandleFunc("/auth/admlogin", Admins.ShowSignupForm)
	r.HandleFunc("/auth/admlogin/post", Admins.SubmitSignupForm)

	http.ListenAndServe(":8080", CSRF(r))
}