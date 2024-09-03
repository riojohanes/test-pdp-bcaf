package services

import (
	"errors"
	"sync"
	"test-pdp-bcaf/models"
)

// interface
type DocumentService interface {
	GetAllDocuments() []models.Document
	GetDocumentByID(id int) (*models.Document, error)
	UpdateDocumentStatus(id int, status string) (*models.Document, error)
	AddDocument(doc models.Document)
	DeleteDocument(id int) error
}

type documentServiceImpl struct {
	mu        sync.Mutex
	documents map[int]models.Document
}

// creates a new instance of DocumentService
func NewDocumentService() DocumentService {
	return &documentServiceImpl{
		documents: map[int]models.Document{
			1: {ID: 1, Name: "Document 1", Status: "pending"},
			2: {ID: 2, Name: "Document 2", Status: "pending"},
			3: {ID: 3, Name: "Document 3", Status: "pending"},
		},
	}
}

// Get all documents
func (s *documentServiceImpl) GetAllDocuments() []models.Document {
	s.mu.Lock()
	defer s.mu.Unlock()

	docs := make([]models.Document, 0, len(s.documents))
	for _, doc := range s.documents {
		docs = append(docs, doc)
	}
	return docs
}

// retrieves a document by its ID
func (s *documentServiceImpl) GetDocumentByID(id int) (*models.Document, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	doc, exists := s.documents[id]
	if !exists {
		return nil, errors.New("document not found")
	}
	return &doc, nil
}

// updates the status of a document
func (s *documentServiceImpl) UpdateDocumentStatus(id int, status string) (*models.Document, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	doc, exists := s.documents[id]
	if !exists {
		return nil, errors.New("document not found")
	}

	doc.Status = status
	s.documents[id] = doc
	return &doc, nil
}

// adds a new document
func (s *documentServiceImpl) AddDocument(doc models.Document) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.documents[doc.ID] = doc
}

// deletes a document by its ID
func (s *documentServiceImpl) DeleteDocument(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.documents[id]; !exists {
		return errors.New("document not found")
	}

	delete(s.documents, id)
	return nil
}
