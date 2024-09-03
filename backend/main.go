package main

import (
	"net/http"

	"test-pdp-bcaf/handlers"
	"test-pdp-bcaf/services"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize service
	documentService := services.NewDocumentService()

	// Initialize handler
	documentHandler := handlers.NewDocumentHandler(documentService)

	// Create a new router
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/documents", documentHandler.GetDocuments).Methods("GET")
	r.HandleFunc("/documents/{id}/approve", documentHandler.ApproveDocument).Methods("POST")
	r.HandleFunc("/documents/{id}/reject", documentHandler.RejectDocument).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", r)
}
