package common

import "net/http"

func AdminsAuthorize(w http.ResponseWriter , r* http.Request , next http.HandlerFunc)  {
	token :=  ""

	if token == "55555" {
		next(w, r)
	} else {
		//http.Redirect(w ,r , "/QAdmin/login" ,http.StatusFound)

		Redirect(w , "/admlogin" ,http.StatusFound)
	}
}
