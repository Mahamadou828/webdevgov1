package main

import (
	"fmt"
	"net/http"
)

/**
* @description it's the function who is call each type someone access our api
 */
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome class</h1>")
	fmt.Println(*r)
}

func main() {
	http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":3000", nil)
}
