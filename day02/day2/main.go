package main

import(
	"fmt"
	"net/http"
	"day2/gee"
)
func handler(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path{
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w,"404 NOT FOUND, url = %q",req.URL)
	}
}
func main(){
	r := gee.New()
	r.GET("/",handler)
	r.GET("/hello",handler)
	r.Run(":9999")
}
