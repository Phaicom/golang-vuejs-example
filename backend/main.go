package main

import (
	"log"
	"net/http"
	"os"

	"github.com/phaicom/golang-vuejs-example/backend/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/phaicom/golang-vuejs-example/backend/endpoints"
	"github.com/phaicom/golang-vuejs-example/backend/middleware"
)

func main() {
	r := mux.NewRouter()

	// Route
	r.HandleFunc("/api/post", endpoints.PostHandler).Methods("GET", "POST")
	r.HandleFunc("/api/reset", endpoints.ResetHandler).Methods("POST")

	// Middleware
	handler := handlers.CORS()(middleware.SecureHeaders(handlers.RecoveryHandler()(handlers.LoggingHandler(os.Stdout, r))))

	// Initial mocking data
	models.ResetJSON()

	log.Fatal(http.ListenAndServe(":8090", handler))
}
