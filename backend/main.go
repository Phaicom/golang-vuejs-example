package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/phaicom/golang-vuejs-example/backend/endpoints"
	"github.com/phaicom/golang-vuejs-example/backend/middleware"
)

func main() {
	r := mux.NewRouter()

	// Route
	r.HandleFunc("/api/post", endpoints.PostHandler).Methods("GET", "POST")

	// Middleware
	handler := handlers.CORS()(middleware.SecureHeaders(handlers.RecoveryHandler()(handlers.LoggingHandler(os.Stdout, r))))

	log.Fatal(http.ListenAndServe(":8090", handler))
}
