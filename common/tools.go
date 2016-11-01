package common

import (
	"net/http"
	"github.com/gorilla/csrf"
)


var (
	crsfKey = []byte("keep-it-secret-keep-it-safe-----")
	CSRF = csrf.Protect(crsfKey)
)




func Redirect(w http.ResponseWriter , url string ,code int )  {
	//output.Context.ResponseWriter.Header().Set(key, val)

	//ctx.Output.Header("Location", localurl)
	//ctx.ResponseWriter.WriteHeader(status)
	w.Header().Set("Location" , url)
	w.WriteHeader(code)
}

