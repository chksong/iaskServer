package Web

import (
	"html/template"
	"net/http"
)

var templates map[string]*template.Template

func Init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("templates/index.html" ,"templates/base.html"))


}

func renderTemplate(w http.ResponseWriter, name string ,template string , viewModel interface{})  {
	tmpl , ok := templates[name]
	if !ok {
		http.Error(w , "The template dost not exist." , http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template ,viewModel)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}