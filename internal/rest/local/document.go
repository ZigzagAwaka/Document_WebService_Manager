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
	return model.Document{}, model.ErrDocNotFound(id)
}

func (s documentService) AddNewElement(document model.Document) error {
	if _, err := s.GetElement(document.ID); err != nil {
		Basic_documents = append(Basic_documents, document)
		return nil
	}
	return model.ErrDocAlreadyExists(document.ID)
}

func (s documentService) DeleteElement(id int) error {
	for i, document := range Basic_documents {
		if id == document.ID {
			Basic_documents = append(Basic_documents[:i], Basic_documents[i+1:]...)
			return nil
		}
	}
	return model.ErrDocNotFound(id)
}
