package main

import (
	"github.com/kingreatwill/go_study/GolangTraining/47_templates/04_template_csv-parse/parse"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// parse csv
	records := parse.Parse("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// function
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// execute template
		err = tpl.Execute(res, records)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// create server
	http.ListenAndServe(":9000", nil)
}
