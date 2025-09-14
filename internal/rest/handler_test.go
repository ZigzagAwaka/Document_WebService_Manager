package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
	"github.com/ZigzagAwaka/Document_WebService_Manager/service"
	"github.com/gin-gonic/gin"
)

func CreateTestingDoc(id int) model.Document {
	return model.Document{ID: id, Title: "Doc" + fmt.Sprint(id), Description: "Content0"}
}

// mockService implements the service.Service interface for testing
type mockService struct {
	elements []model.Document
}

func (mockService) KeyWord() string {
	return "documents"
}

func (m mockService) GetAllElements() []model.Document {
	return m.elements
}

func (m mockService) GetElement(id int) (model.Document, error) {
	for _, e := range m.elements {
		if e.ID == id {
			return e, nil
		}
	}
	return model.Document{}, model.ErrDocNotFound(id)
}

func (m *mockService) AddNewElement(doc model.Document) error {
	if _, err := m.GetElement(doc.ID); err != nil {
		m.elements = append(m.elements, doc)
		return nil
	}
	return model.ErrDocAlreadyExists(doc.ID)
}

func (m *mockService) DeleteElement(id int) error {
	for i, e := range m.elements {
		if e.ID == id {
			m.elements = append(m.elements[:i], m.elements[i+1:]...)
			return nil
		}
	}
	return model.ErrDocNotFound(id)
}

func setupRouter(service service.Service) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewHandler(service)
	handler.Init(router)
	return router
}

//////////////////////////////////////////////////////////
// Test cases for documentHandler
//////////////////////////////////////////////////////////

func TestGetAllElements(t *testing.T) {
	service := &mockService{
		elements: []model.Document{CreateTestingDoc(1), CreateTestingDoc(2)},
	}
	router := setupRouter(service)

	req, _ := http.NewRequest("GET", "/documents", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, r.Code)
	}
	for i := 1; i <= 2; i++ {
		if !bytes.Contains(r.Body.Bytes(), fmt.Appendf(nil, "Doc%d", i)) {
			t.Errorf("Expected body to contain 'Doc%d', got %s", i, r.Body.String())
		}
	}
}

func TestGetElement_Success(t *testing.T) {
	service := &mockService{
		elements: []model.Document{CreateTestingDoc(8)},
	}
	router := setupRouter(service)

	req, _ := http.NewRequest("GET", "/documents/8", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Doc8")) {
		t.Errorf("Expected body to contain 'Doc8', got %s", r.Body.String())
	}
}

func TestGetElement_NotFound(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	req, _ := http.NewRequest("GET", "/documents/99", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, r.Code)
	}
	err := model.ErrDocNotFound(99).Error()
	if !bytes.Contains(r.Body.Bytes(), []byte(err)) {
		t.Errorf("Expected body to contain %s but got %s", err, r.Body.String())
	}
}

func TestGetElement_InvalidID(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	req, _ := http.NewRequest("GET", "/documents/abc", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Invalid request body")) {
		t.Errorf("Expected body to contain an error but got %s", r.Body.String())
	}
}

func TestAddNewElement_Success(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	body := []byte(`{"ID":3,"Title":"Doc3","Description":"Content0"}`)
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Doc3")) {
		t.Errorf("Expected body to contain 'Doc3', got %s", r.Body.String())
	}
	if len(service.elements) != 1 {
		t.Errorf("Expected size of elements to be 1 after addition, got %d", len(service.elements))
	}
}

func TestAddNewElement_BadRequest(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	body := []byte(`{invalid json}`)
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Invalid request body")) {
		t.Errorf("Expected body to contain an error but got %s", r.Body.String())
	}
}

func TestAddNewElement_Conflict(t *testing.T) {
	service := &mockService{
		elements: []model.Document{CreateTestingDoc(4)},
	}
	router := setupRouter(service)

	body := []byte(`{"ID":4,"Title":"Doc4","Description":"Content0"}`)
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusConflict {
		t.Errorf("Expected status %d, got %d", http.StatusConflict, r.Code)
	}
	err := model.ErrDocAlreadyExists(4).Error()
	if !bytes.Contains(r.Body.Bytes(), []byte(err)) {
		t.Errorf("Expected body to contain %s but got %s", err, r.Body.String())
	}
}

func TestDeleteElement_Success(t *testing.T) {
	service := &mockService{
		elements: []model.Document{CreateTestingDoc(5)},
	}
	router := setupRouter(service)

	req, _ := http.NewRequest("DELETE", "/documents/5", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusAccepted {
		t.Errorf("Expected status %d, got %d", http.StatusAccepted, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Element successfully deleted")) {
		t.Errorf("Expected body to contain 'Element successfully deleted' but got %s", r.Body.String())
	}
	if len(service.elements) != 0 {
		t.Errorf("Expected size of elements to be 0 after deletion, got %d", len(service.elements))
	}
}

func TestDeleteElement_NotFound(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	req, _ := http.NewRequest("DELETE", "/documents/99", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, r.Code)
	}
	err := model.ErrDocNotFound(99).Error()
	if !bytes.Contains(r.Body.Bytes(), []byte(err)) {
		t.Errorf("Expected body to contain %s but got %s", err, r.Body.String())
	}
}

func TestDeleteElement_InvalidID(t *testing.T) {
	service := &mockService{}
	router := setupRouter(service)

	req, _ := http.NewRequest("DELETE", "/documents/abc", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, r.Code)
	}
	if !bytes.Contains(r.Body.Bytes(), []byte("Invalid request body")) {
		t.Errorf("Expected body to contain an error but got %s", r.Body.String())
	}
}
