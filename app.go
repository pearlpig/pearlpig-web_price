package main

import (
	"fmt"
	"net/http"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/julienschmidt/httprouter"
)

func main() {
	host := "127.0.0.1"
	port := "8080"
	// Set the routes for the web application.
	mux := httprouter.New()

	// Listen to index page
	mux.GET("/", indexHandler)

	/* Create the logger for the web application. */
	l := log.New()

	n := negroni.New()
	n.Use(negronilogrus.NewMiddlewareFromLogger(l, "web"))
	n.UseHandler(mux)
	// Set the parameters for a HTTP server
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: n,
	}

	// Listen requests
	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	// log.Fatal()
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// t, err := template.ParseFiles("views/layout.html", "views/head.html", "views/index.html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`welcome`))
}
