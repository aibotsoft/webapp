package main

import (
	"fmt"
	"net/http"
)

func handlerFuck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")

}
func main() {
	http.HandleFunc("/", handlerFuck)
	fmt.Printf("http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Print(err)
	}
}
