package main

import (
	"fmt"
	"github.com/aibotsoft/webapp/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	must(homeView.Render(w, nil))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	must(contactView.Render(w, nil))
}
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	must(signupView.Render(w, nil))

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
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	//fmt.Printf("http://localhost:3000")
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/faq", faq)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
