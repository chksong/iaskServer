package Web

import "net/http"
import (
	"html/template"
	"time"
	"fmt"
)

func Test (w http.ResponseWriter , r* http.Request)  {
	//renderTemplate(w , "index" ,"base" ,nil)

	tmpl := template.Must(template.ParseFiles("templates/test.html" ,"templates/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}

func Index (w http.ResponseWriter , r* http.Request)  {
	//renderTemplate(w , "index" ,"base" ,nil)


	cookie ,_ := r.Cookie("uid")
	fmt.Println("readCookie===" ,cookie )

	COOKIE_MAX_MAX_AGE     := time.Hour * 24 / time.Second   // 单位：秒。
	maxAge := int(COOKIE_MAX_MAX_AGE)
	uid:="10"

	uid_cookie:=&http.Cookie{
		Name:   "uid",
		Value:    uid,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   maxAge ,
	}

	http.SetCookie(w,uid_cookie)




	tmpl := template.Must(template.ParseFiles("templates/Web/index.html" ,"templates/Web/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}