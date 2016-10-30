package common

import "net/http"

func Redirect(w http.ResponseWriter , url string ,code int )  {
	//output.Context.ResponseWriter.Header().Set(key, val)

	//ctx.Output.Header("Location", localurl)
	//ctx.ResponseWriter.WriteHeader(status)
	w.Header().Set("Location" , url)
	w.WriteHeader(code)
}

