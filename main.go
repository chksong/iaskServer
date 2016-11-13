package main
import (

	"net/http"
	"fmt"
	"iaskServer/routers"
	"log"

	"iaskServer/common"
	"github.com/codegangsta/negroni"

	"iaskServer/csrf"
)

func hello(w http.ResponseWriter , r* http.Request)  {
	fmt.Fprintf(w, "hello")
}



func main() {
	//calls startup logic
	common.StartUp()

	//出事话路由
	router := routers.InitRouters()


	n := negroni.Classic()
	n.UseHandler(router)

	log.Println("Listening ........")
	//server := &http.Server{
	//	Addr:common.AppConfig.Server,
	//	Handler:n ,
	//}
	//生成32字符
	//openssl rand -base64 32 |head  -c 32
	crsfKey := []byte("tTKCH86aTVEphRM5LrD6Ps/usAbLAQQr")
	CSRF := csrf.Protect(crsfKey,csrf.Secure(false))

	//添加csrf 保护
	http.ListenAndServe(common.AppConfig.Server, CSRF(n))
}
