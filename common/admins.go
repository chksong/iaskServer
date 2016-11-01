package common

import (
	"net/http"
	"time"
	"log"
)


/**
 *    验证管理员权限
 */
func AdminsAuthorize(w http.ResponseWriter , r* http.Request , next http.HandlerFunc)  {
	if adminCookie ,err := r.Cookie("iaskiAdminToken") ; err == nil {
		if adminCookie.Value == "99999" {
			next(w, r)
		} else {
			//http.Redirect(w ,r , "/QAdmin/login" ,http.StatusFound)

			Redirect(w , "/auth/admlogin" ,http.StatusFound)
		}
	} else {
		Redirect(w , "/auth/admlogin" ,http.StatusFound)
	}
}

/**
 *   日志帮主函数
 */
func LoggingHander(w http.ResponseWriter , r* http.Request , next http.HandlerFunc) {

	start := time.Now()
	log.Printf("START %s %s" , r.Method ,r.URL.Path)
	next(w ,r )
	log.Printf("Completed %s in %v" , r.URL.Path , time.Since(start))
}