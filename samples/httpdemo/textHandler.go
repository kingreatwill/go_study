
package main

import (
	"fmt"
	"net/http"
)

type textHandler struct {
	h http.HandlerFunc
}

func (t textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set the content type
	w.Header().Set("Content-Type", "text/plain")
	// Then call ServeHTTP in the decorated handler.
	t.h(w, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	http.Handle("/hello", textHandler{helloHandler})
	http.ListenAndServe(":8080", nil)
}
