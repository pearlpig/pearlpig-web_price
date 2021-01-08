package main

import (
	"fmt"
	"html/template"
	"net/http"
	"pearlpig/web_price/crawler"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// t, err := template.ParseFiles("views/layout.html", "views/head.html", "views/index.html")
	var tmpl = template.Must(template.ParseFiles("views/index.html", "views/head.html", "views/body.html"))

	tmpl.ExecuteTemplate(w, "index", struct {
		Title string
	}{
		"比價網站",
	})
}

func updateResultHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Parse the form failed")
	}

	request := r.FormValue("request")
	result := crawler.App(request)

	// request := r.FormValue("search")
	// t1 := time.Now() // get current time
	// result := crawler.CrawlItems(r.Form)
	// elapsed := time.Since(t1)
	// fmt.Println("App elapsed: ", elapsed)
	w.Write(result)
	// fmt.Fprintf(w, result)
}
