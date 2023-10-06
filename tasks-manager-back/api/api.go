package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func SetupHandlers() {
	Router = httprouter.New()

	AddUserHandlers()

	// Add Webpage
	Router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./public/index.html")
	})
	Router.RedirectTrailingSlash = true
	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", Router))
}
