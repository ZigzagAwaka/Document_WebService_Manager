package service

import "github.com/ZigzagAwaka/Document_WebService_Manager/model"

// Common interface for all document services
type Service interface {
	// Return the service keyword
	KeyWord() string

	// Get all documents
	GetAllElements() []model.Document

	// Get a specific document by ID
	GetElement(id int) (model.Document, error)

	// Add a new document
	AddNewElement(document model.Document) error
}
