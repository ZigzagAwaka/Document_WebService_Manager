package local

import (
	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
	"github.com/ZigzagAwaka/Document_WebService_Manager/service"
)

type documentService struct{}

// Create a document service as a service to manage locally stored documents
func NewLocalDocumentService() service.Service {
	return &documentService{}
}

func (documentService) KeyWord() string {
	return "documents"
}

func (documentService) GetAllElements() []model.Document {
	return Basic_documents
}

func (documentService) GetElement(id int) (model.Document, error) {
	for _, document := range Basic_documents {
		if id == document.ID {
			return document, nil
		}
	}
	return model.Document{}, nil
	//context.String(http.StatusNotFound, "Document with ID %v not found", id)
}

func (documentService) AddNewElement(document model.Document) {
	Basic_documents = append(Basic_documents, document)
}
