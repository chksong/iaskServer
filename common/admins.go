package common

import (
	"net/http"
	"time"
	"log"
//	"fmt"
)


/**
 *    验证管理员权限, 不需要添加跳转
 */
func AdminsAuthorize(w http.ResponseWriter , r* http.Request , next http.HandlerFunc)  {
	//if adminCookie ,err := r.Cookie("iaskiAdminToken") ; err == nil {
	//	if adminCookie.Value == "99999" {
	//		next(w, r)
	//	} else {
	//		http.Redirect(w ,r , "/QAdmin/login2" ,http.StatusMovedPermanently)
	//
	//		//Redirect(w , "/auth/admlogin" ,http.StatusFound)
	//		//fmt.Fprintln(w, "前先登录")
	//
	//	}
	//} else {
	//	// 没有验证过，所以验证登录
	//	//Redirect(w , "/auth/admlogin" ,http.StatusFound)
	//}
	next(w,r)
}


/**
 *  验证Crsf
 */
func CrsfAuthorize(w http.ResponseWriter , r* http.Request , next http.HandlerFunc)  {

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