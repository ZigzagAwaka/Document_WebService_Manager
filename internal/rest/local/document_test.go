package local

import (
	"testing"

	"github.com/ZigzagAwaka/Document_WebService_Manager/model"
)

func CreateTestingDoc(id int) model.Document {
	return model.Document{ID: id, Title: "Testing document", Description: "Content0"}
}

func TestGetAllInitial(t *testing.T) {
	service := NewLocalDocumentService()
	docs := service.GetAllElements()
	if len(docs) != len(Basic_documents) {
		t.Errorf("GetAllElements() has returned %v elements but %v elements were expected at the initial state", len(docs), len(Basic_documents))
	}
}

func TestGetAllAdded(t *testing.T) {
	nb := 5
	service := NewLocalDocumentService()
	initSize := len(service.GetAllElements())
	for i := range nb {
		service.AddNewElement(CreateTestingDoc(100 + i))
	}
	docs := service.GetAllElements()
	if len(docs) != initSize+nb {
		t.Errorf("GetAllElements() has returned %v elements but %v elements were expected after adding %v more elements", len(docs), initSize+nb, nb)
	}
	for i := range nb {
		service.DeleteElement(100 + i)
	}
}

func TestGetAllRemoved(t *testing.T) {
	nb := 3
	service := NewLocalDocumentService()
	for i := 1; i < nb+1; i++ {
		service.AddNewElement(CreateTestingDoc(100 + i))
	}
	initSize := len(service.GetAllElements())
	for i := 1; i < nb+1; i++ {
		service.DeleteElement(100 + i)
	}
	docs := service.GetAllElements()
	if len(docs) != initSize-nb {
		t.Errorf("GetAllElements() has returned %v elements but %v elements were expected after deleting %v elements", len(docs), initSize+nb, nb)
	}
}

func TestGetElementFound(t *testing.T) {
	service := NewLocalDocumentService()
	for _, doc := range Basic_documents {
		retDoc, err := service.GetElement(doc.ID)
		if retDoc != doc || err != nil {
			t.Errorf("GetElement(%v) returned %v but %v was expected, %v", doc.ID, retDoc, doc, err)
			return
		}
	}
}

func TestGetElementNotFound(t *testing.T) {
	service := NewLocalDocumentService()
	_, err := service.GetElement(1000)
	if err == nil {
		t.Errorf("GetElement(1000) error statement was [%v] but an error was expected", err)
	}
}

func TestGetElementRecentlyAdded(t *testing.T) {
	service := NewLocalDocumentService()
	newDoc := CreateTestingDoc(1000)
	retDoc, err := service.GetElement(newDoc.ID)
	if retDoc == newDoc || err == nil {
		t.Errorf("GetElement(%v) returned %v but %v was expected, %v", newDoc.ID, retDoc, newDoc, err)
	}
	service.AddNewElement(newDoc)
	retDoc, err = service.GetElement(newDoc.ID)
	if retDoc != newDoc || err != nil {
		t.Errorf("GetElement(%v) returned %v but %v was expected, %v", newDoc.ID, retDoc, newDoc, err)
	}
	service.DeleteElement(newDoc.ID)
}

func TestAddNewElementSuccess(t *testing.T) {
	service := NewLocalDocumentService()
	newDoc := CreateTestingDoc(1000)
	err := service.AddNewElement(newDoc)
	if err != nil {
		t.Errorf("AddNewElement(%v) error statement was [%v] but no error was expected", newDoc, err)
	}
	retDoc, err := service.GetElement(newDoc.ID)
	if retDoc != newDoc || err != nil {
		t.Errorf("GetElement(%v) returned %v but %v was expected, %v", newDoc.ID, retDoc, newDoc, err)
	}
	service.DeleteElement(newDoc.ID)
}

func TestAddNewElementFailure(t *testing.T) {
	service := NewLocalDocumentService()
	existingDoc := Basic_documents[0]
	err := service.AddNewElement(existingDoc)
	if err == nil {
		t.Errorf("AddNewElement(%v) error statement was [%v] but an error was expected after adding an already existing document", existingDoc, err)
	}
}

func TestDeleteElementSuccess(t *testing.T) {
	service := NewLocalDocumentService()
	newDoc := CreateTestingDoc(1000)
	service.AddNewElement(newDoc)
	err := service.DeleteElement(newDoc.ID)
	if err != nil {
		t.Errorf("DeleteElement(%v) error statement was [%v] but no error was expected", newDoc.ID, err)
	}
	_, err = service.GetElement(newDoc.ID)
	if err == nil {
		t.Errorf("GetElement(%v) error statement was [%v] but an error was expected after getting a deleted document", newDoc.ID, err)
	}
}

func TestDeleteElementFailure(t *testing.T) {
	service := NewLocalDocumentService()
	err := service.DeleteElement(1000)
	if err == nil {
		t.Errorf("DeleteElement(1000) error statement was [%v] but an error was expected when deleting a non-existing document", err)
	}
}
