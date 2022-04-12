package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kellanjacobs/icanhazquotes/config"
	"github.com/kellanjacobs/icanhazquotes/db"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	c, _ := config.LoadConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.App.Hostname, c.App.Port), mux))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	q := db.GetRandomQuote()
	err := tpl.ExecuteTemplate(w, "index.gohtml", q)
	if err != nil {
		log.Fatalln(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	dbgood := db.CheckDB()
	if dbgood {
		w.Write([]byte("200 - OK!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
}
