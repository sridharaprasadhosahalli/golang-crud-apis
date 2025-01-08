package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
    "bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestGetItemsHandler(t *testing.T) {
    r := SetUpRouter()
    r.GET("/items", getItems)
    req, _ := http.NewRequest("GET", "/items", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var items []item
    json.Unmarshal(w.Body.Bytes(), &items)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, items)
}

func TestGetItemByIdHandler(t *testing.T) {
    r := SetUpRouter()
    r.GET("/items/:id", getItemByID)
    req, _ := http.NewRequest("GET", "/items/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var items item
    json.Unmarshal(w.Body.Bytes(), &items)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, items)
}

func TestDeleteItemByIdHandler(t *testing.T) {
    r := SetUpRouter()
    r.DELETE("/items/:id",deleteItem)
    req, _ := http.NewRequest("DELETE", "/items/3", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(items))
}

func TestAddItemHandler(t *testing.T) {
    r := SetUpRouter()
    r.POST("/items", addItem)
    itemm := item{
        ID: "4",
        Name: "Orange",
        Type: "Fruit",
        Price: 60,
    }
    jsonValue, _ := json.Marshal(itemm)
    req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateItemHandler(t *testing.T) {
    r := SetUpRouter()
    r.PUT("/items/:id",updateItem)
    itemm := item{
        ID: "1",
        Name: "Orange",
        Type: "Fruit",
        Price: 60,
    }
    jsonValue, _ := json.Marshal(itemm)
    req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
	
	reqNotFound, _ := http.NewRequest("PUT", "/items/12", bytes.NewBuffer(jsonValue))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, reqNotFound)
    assert.Equal(t, http.StatusNotFound, w.Code)
}