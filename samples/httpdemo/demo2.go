
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"

)

func main() {
	http.HandleFunc("/body", bodyHandler)
	http.HandleFunc("/param", paramHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "could not read body: %v", err)
		return
	}
	name := string(b)
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
