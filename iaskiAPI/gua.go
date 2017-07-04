package iaskiAPI

import (
	"net/http"
	"github.com/unrolled/render"
)



var (
	formatter = render.New(render.Options{IndentJSON:true})
)


func TestJson(w http.ResponseWriter , r * http.Request)  {

	formatter.JSON(w ,http.StatusOK ,
		struct{Test string }{"THis is a test "})
}