package local

import (
	"net/http"
	"strconv"

	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
	"github.com/ZigzagAwaka/Document_WebService_Manager/service"
	"github.com/gin-gonic/gin"
)

// Service for documents
type documentService struct{}

// Create a new document service as a service
func NewDocumentService() service.Service {
	return &documentService{}
}

func (documentService) GetAllElements(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, model.Basic_documents)
}

func (documentService) GetElement(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}

	for _, document := range model.Basic_documents {
		if id == document.ID {
			context.IndentedJSON(http.StatusOK, document)
			return
		}
	}
	context.String(http.StatusNotFound, "Document with ID %v not found", id)
}

func (documentService) AddNewElement(context *gin.Context) {
	var newDocument model.Document

	if err := context.BindJSON(&newDocument); err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}
	model.Basic_documents = append(model.Basic_documents, newDocument)
	context.IndentedJSON(http.StatusCreated, newDocument)
}
