

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// A sample JSON equivalent
/*
{
	"name": "gopher",
	"age_years": 5
}
*/

// Person contains all the information I can imagine about a person right now.
type Person struct {
	Name     string `json:"name"`
	AgeYears int    `json:"age_years"`
}

func encode() {
	p := Person{"gopher", 5}

	// create an encoder that will write on the standard output.
	enc := json.NewEncoder(os.Stdout)
	// use the encoder to encode p, which could fail.
	err := enc.Encode(p)
	// if it failed, log the error and stop execution.
	if err != nil {
		log.Fatal(err)
	}
}

func decode() {
	// create an empty Person value.
	var p Person

	// create a decoder reading from the standard input.
	dec := json.NewDecoder(os.Stdin)
	// use the decoder to decode a value into p.
	err := dec.Decode(&p)
	// if it failed, log the error and stop execution.
	if err != nil {
		log.Fatal(err)
	}
	// otherwise log what we decoded.
	fmt.Printf("decoded: %#v\n", p)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{"gopher", 5}

	// set the Content-Type header.
	w.Header().Set("Content-Type", "application/json")

	// encode p to the output.
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	if err != nil {
		// if encoding fails, create an error page with code 500.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var p Person

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name is %v and age is %v", p.Name, p.AgeYears)
}

func main() {
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
