package main

import (
	"fmt"
	"net/http"
)

/**
* @description it's the function who is call each type someone access our api
 */
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	path := r.URL.RequestURI()

	switch path {
	case "/support":
		fmt.Fprint(w, "To get in touch send a email to <a href=\"mailto:opmadou@gmail.com\"/>Support me!</a>")
	default:
		fmt.Fprint(w, "<h1>Welcome to my awesome class</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Up And Serve!")
	http.ListenAndServe(":3000", nil)
}
