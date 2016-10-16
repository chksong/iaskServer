package Admins

import (
	"html/template"
	"net/http"
)

func Index (w http.ResponseWriter , r* http.Request)  {
	//renderTemplate(w , "index" ,"base" ,nil)

	tmpl := template.Must(template.ParseFiles("templates/Admins/index.html" ,"templates/Admins/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}

func Login (w http.ResponseWriter , r* http.Request)  {
	tmpl := template.Must(template.ParseFiles("templates/Admins/login.html" ,"templates/Admins/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}