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


func main__() {
	//	Routes:
	r := mux.NewRouter().StrictSlash(false)

	//	Root 'home' route
	r.HandleFunc("/", HomeHandler)

	//	To login/logout/signup:
	//	/auth/login
	//	/auth/logout
	//	/auth/signup
	auth := r.PathPrefix("/auth").Subrouter()
	auth.Path("/login").HandlerFunc(LoginHandler)
	auth.Path("/logout").HandlerFunc(LogoutHandler)
	auth.Path("/signup").HandlerFunc(SignupHandler)

	// Posts collection
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	//	Accounts
	acctBase := mux.NewRouter()
	r.PathPrefix("/account").Handler(negroni.New(
	negroni.NewRecovery(),
	negroni.HandlerFunc(MyMiddleware),
	negroni.NewLogger(),
	negroni.Wrap(acctBase),
	))

	acct := acctBase.PathPrefix("/account").Subrouter()
	acct.Path("/profile").HandlerFunc(ProfileHandler)

	// Posts singular
	post := r.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
}

