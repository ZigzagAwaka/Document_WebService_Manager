package rest

import (
	"net/http"
	"strconv"

	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
	"github.com/ZigzagAwaka/Document_WebService_Manager/service"
	"github.com/gin-gonic/gin"
)

// Struct implementing a document service to manage documents
type documentRepository struct {
	service service.Service
}

// Create a document repository with the given management service
func NewRepository(service service.Service) *documentRepository {
	return &documentRepository{service}
}

// Initialize the given Gin router with the repository endpoints
func (r *documentRepository) Init(router *gin.Engine) {
	key := r.service.KeyWord()
	router.GET("/"+key, r.getAllElements)
	router.GET("/"+key+"/:id", r.getElement)
	router.POST("/"+key, r.addNewElement)
}

func (r *documentRepository) getAllElements(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, r.service.GetAllElements())
}

func (r *documentRepository) getElement(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}
	element, err := r.service.GetElement(id)
	if err != nil {
		context.String(http.StatusNotFound, "Error when getting the element: %v", err)
		return
	}
	context.IndentedJSON(http.StatusOK, element)
}

func (r *documentRepository) addNewElement(context *gin.Context) {
	var newElement model.Document
	if err := context.BindJSON(&newElement); err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}
	r.service.AddNewElement(newElement)
	context.IndentedJSON(http.StatusCreated, newElement)
}
