package service

import (
	"github.com/gin-gonic/gin"
)

// Common interface for all services
type Service interface {
	GetAllElements(context *gin.Context)
}
