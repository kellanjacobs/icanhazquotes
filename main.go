package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/kellanjacobs/icanhazquotes/quotes"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	q := quotes.GetRandomQuote()
	err := tpl.ExecuteTemplate(w, "index.gohtml", q)
	if err != nil {
		log.Fatalln(err)
	}
}
