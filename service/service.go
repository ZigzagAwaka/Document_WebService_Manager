package service

import (
	"github.com/gin-gonic/gin"
)

// Common interface for all services
type Service interface {
	// Get all elements of the current service
	GetAllElements(context *gin.Context)

	// Add a new element to the current service
	AddNewElement(context *gin.Context)
}
