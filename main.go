package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Home page!</h1>")

}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Contact page</h1>")
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Frequently Asked Questions Page</h1>")
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprint(w, "<h1>Not found page 404</h1>")
}

func main() {
	//fmt.Printf("http://localhost:3000")
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	log.Fatal(http.ListenAndServe(":3000", r))
}
