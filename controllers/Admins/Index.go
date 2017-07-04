package Admins

import (
	"html/template"
	"net/http"
//	"iaskServer/common"
	"log"
	"github.com/gorilla/csrf"
	"fmt"
	"iaskServer/common"
)


func Index (w http.ResponseWriter , r* http.Request)  {
	//renderTemplate(w , "index" ,"base" ,nil)

	cookie := http.Cookie{
		Name:"adminCookie" ,
		Value: "99999" ,
		Path: "/",
		Secure:false  ,
		HttpOnly:true ,
		MaxAge: 3600 ,
	}
	http.SetCookie(w , &cookie)



	tmpl := template.Must(template.ParseFiles("templates/Admins/index.html" ,"templates/Admins/base.html"))
	err := tmpl.ExecuteTemplate(w, "base" ,nil)
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
}

func ShowSignupForm(w http.ResponseWriter , r* http.Request)  {

	cookie := http.Cookie{
		Name:"testCookie" ,
		Value: "99999" ,
		Path: "/token/",
		Secure:false  ,
		HttpOnly:false ,
		MaxAge: 3600 ,
	}

	http.SetCookie(w , &cookie)



	cookie2, err_ := r.Cookie("testCookie")
	fmt.Println("testCookie=======", cookie2 ,"err_==" ,err_)




	tmpl := template.Must(template.ParseFiles("templates/Admins/login.html" ,"templates/Admins/base.html"))

	locals := make(map[string]interface{})
	//locals["title"]  = "test"
	locals[csrf.TemplateTag] = csrf.TemplateField(r)


	w.Header().Set("X-CSRF-Token", csrf.Token(r))


	err := tmpl.ExecuteTemplate(w, "base" ,locals)

	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}


}

/**
 *   检查admin权限
 */
func SubmitSignupForm(w http.ResponseWriter , r* http.Request)  {
	log.Printf("[CheckAdmin]   ......") ;

	fmt.Println("[ShowSignupForm] username: ", r.Form["username"])
	fmt.Println("[ShowSignupForm] password: ", r.Form["passwd"])

	cookie := &http.Cookie{
		Name:"iaskiAdminToken" ,
		Value: "99999" ,
		Path: "/",
		Secure:false  ,
		HttpOnly:true ,
		MaxAge: 3600 ,
		Domain:"iaski.com" ,
	}




	http.SetCookie(w , cookie)
	common.Redirect(w , "/QAdmins/index" ,http.StatusMovedPermanently)
}