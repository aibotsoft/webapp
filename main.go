package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("favicon requested")
}
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	//_, _ = fmt.Fprint(w, r.URL.Path)
	fmt.Println(r.URL)
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Welcome to my awesome site change!</h1>")
	} else if r.URL.Path == "/fuck" {
		_, _ = fmt.Fprint(w, "<h1>Fuck off</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "<h1>I do not have that page</h1>")
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	//mux := &http.ServeMux{}
	router := httprouter.New()
	//router.GET("/", handlerFunc)
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	//mux.HandleFunc("/", handlerFunc)
	//fmt.Printf("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
