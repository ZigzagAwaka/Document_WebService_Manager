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

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", GetHomePage)
	router.GET("/documents", service.GetAllElements)

	log.Println("Service initialized, listening on http://" + serverAddress + ".")

	// Start the server (in local)
	router.Run(serverAddress)
}
