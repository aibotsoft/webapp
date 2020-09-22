package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	_, _ = fmt.Fprint(w, r.URL.Path)
	fmt.Println(r.URL)
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Welcome to my awesome site change!</h1>")
	} else if r.URL.Path == "/fuck" {
		_, _ = fmt.Fprint(w, "<h1>Fuck off</h1>")
	}

}
func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Printf("http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Print(err)
	}
}
