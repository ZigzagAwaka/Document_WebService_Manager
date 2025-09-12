package rest

import (
	"net/http"
	"strconv"

	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
	"github.com/ZigzagAwaka/Document_WebService_Manager/service"
	"github.com/gin-gonic/gin"
)

// Handler implementing a service to manage documents
type documentHandler struct {
	service service.Service
}

// Create a document handler with the given management service
func NewHandler(service service.Service) *documentHandler {
	return &documentHandler{service}
}

// Initialize the given Gin router with the handler endpoints
func (r *documentHandler) Init(router *gin.Engine) {
	key := r.service.KeyWord()
	router.GET("/"+key, r.getAllElements)
	router.GET("/"+key+"/:id", r.getElement)
	router.POST("/"+key, r.addNewElement)
	router.DELETE("/"+key+"/:id", r.deleteElement)
}

func (r *documentHandler) getAllElements(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, r.service.GetAllElements())
}

func (r *documentHandler) getElement(context *gin.Context) {
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

func (r *documentHandler) addNewElement(context *gin.Context) {
	var newElement model.Document
	if err := context.BindJSON(&newElement); err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}
	if err := r.service.AddNewElement(newElement); err != nil {
		context.String(http.StatusConflict, "Error when adding the element: %v", err)
		return
	}
	context.IndentedJSON(http.StatusCreated, newElement)
}

func (r *documentHandler) deleteElement(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}
	if err := r.service.DeleteElement(id); err != nil {
		context.String(http.StatusNotFound, "Error when deleting the element: %v", err)
		return
	}
	context.String(http.StatusAccepted, "Element successfully deleted")
}
