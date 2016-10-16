package Web

import "net/http"
import "html/template"

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

	tmpl := template.Must(template.ParseFiles("templates/index.html" ,"templates/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}