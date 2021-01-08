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
	// Listen to CSS assets
	mux.ServeFiles("/css/*filepath", http.Dir("public/css"))

	// // Listen to JavaScript assets

	mux.ServeFiles("/js/*filepath", http.Dir("public/js"))

	// Listen to index page
	mux.GET("/", indexHandler)

	// Respond to result
	mux.GET("/result/", updateResultHandler)

	// // Custom 404 page
	// mux.NotFound = http.HandlerFunc(notFound)

	// // Custom 500 page
	// mux.PanicHandler = errorHandler

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

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("views/index.html", "views/notFound.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 	}

// 	w.WriteHeader(http.StatusNotFound) // HTTP 404
// 	t.ExecuteTemplate(w, "layout", "Calculator")
// }

// func errorHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
// 	t, err := template.ParseFiles("views/index.html", "views/error.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 	}

// 	w.WriteHeader(http.StatusInternalServerError) // HTTP 500
// 	t.ExecuteTemplate(w, "layout", "Calculator")
// }
