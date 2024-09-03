package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"test-pdp-bcaf/services"

	"github.com/gorilla/mux"
)

// DocumentHandler handles HTTP requests for documents
type DocumentHandler struct {
	service services.DocumentService
}

// NewDocumentHandler creates a new DocumentHandler
func NewDocumentHandler(service services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		service: service,
	}
}

// GetDocuments
func (h *DocumentHandler) GetDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	documents := h.service.GetAllDocuments()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(documents)
}

// Approved
func (h *DocumentHandler) ApproveDocument(w http.ResponseWriter, r *http.Request) {
	h.updateDocumentStatus(w, r, "approved")
}

// Rejected
func (h *DocumentHandler) RejectDocument(w http.ResponseWriter, r *http.Request) {
	h.updateDocumentStatus(w, r, "rejected")
}

// Updates the document status
func (h *DocumentHandler) updateDocumentStatus(w http.ResponseWriter, r *http.Request, status string) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid document ID", http.StatusBadRequest)
		return
	}

	doc, err := h.service.UpdateDocumentStatus(id, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doc)
}
