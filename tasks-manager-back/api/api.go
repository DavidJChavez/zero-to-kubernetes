package api

import (
	"log"
	"net/http"
)

func SetupHandlers() {

	AddWebPage()

	AddUserHandlers()

	log.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func AddWebPage() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./public/index.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
