package main

import(
	"fmt"
	"log"
	"net/http"
)

type Engine struct{
	Path string
}
func (eng *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request) {
	eng.Path = req.URL.Path;
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
func main(){
	eng := new(Engine)
	log.Fatal(http.ListenAndServe(":8888",eng))
}
