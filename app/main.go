package main

import (
	"log"
	"net/http"

	localRepo "github.com/ZigzagAwaka/Document_WebService_Manager/internal/repository"

	"github.com/gin-gonic/gin"
)

const serverAddress = "localhost:8080"

// Get example home page
func GetHomePage(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to the WebService Manager!")
}

func main() {
	log.SetPrefix("[Document_WebService_Manager] ")
	log.SetFlags(0)
	log.Println("Initializing Service...")

	// Initialize the document service
	service := localRepo.NewDocumentService()

	// Set up the Gin router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", GetHomePage)
	router.GET("/documents", service.GetAllElements)
	router.GET("/documents/:id", service.GetElement)
	router.POST("/documents", service.AddNewElement)

	log.Println("Service initialized, listening on http://" + serverAddress + ".")

	// Start the server (in local)
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
