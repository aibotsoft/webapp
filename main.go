package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "<h1>Fuck off</h1>")

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Welcome to my awesome site change!</h1>")
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Home page!</h1>")

}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Contact page</h1>")
}

func main() {
	//fmt.Printf("http://localhost:3000")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(":3000", r))
}
