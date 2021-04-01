package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/murnux/grocery-store-api/produce_api"
)

// TestDeleteProduceHandler tests the delete produce handler in api.go
func TestDeleteProduceHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	store := produce_api.CreateStore()

	// test produce
	produce := produce_api.Produce{
		Name:  "testing",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 0.0,
	}
	store.ProduceItems = append(store.ProduceItems, produce)

	router := produce_api.CreateRouter(store) // create router object with a store
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/produce/delete?Produce%20Code="+produce.Code, nil)
	router.ServeHTTP(recorder, request)

	// check that the produce item was deleted and that the API returned 200
	if len(store.ProduceItems) != 0 && recorder.Code != 200 {
		t.Errorf("There was an error deleting the produce item")
	}
}

// TestAddProduceHandler tests the add produce handler in api.go
func TestAddProduceHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	store := produce_api.CreateStore()
	// test produce
	produce := produce_api.Produce{
		Name:  "testing",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 0.0,
	}

	// create ProduceList object, which is what the API expects
	produceList := produce_api.ProduceList{}
	produceList.List = make([]produce_api.Produce, 1)
	produceList.List[0] = produce

	router := produce_api.CreateRouter(store)
	recorder := httptest.NewRecorder()
	jsonBytes, _ := json.Marshal(produceList) // convert ProduceList struct to array of bytes for the request
	request, _ := http.NewRequest("POST", "/produce/add", bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json") // send header to tell the API it is getting JSON
	router.ServeHTTP(recorder, request)

	if len(store.ProduceItems) != 1 || recorder.Code != 200 || store.ProduceItems[0].Code != produce.Code {
		t.Errorf("There was an error adding the produce item")
	}
}
