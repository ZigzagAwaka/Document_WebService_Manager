package local

import (
	"net/http"

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

// Get all documents of the current document service
func (documentService) GetAllElements(context *gin.Context) {
	context.JSON(http.StatusOK, model.Basic_documents)
}
