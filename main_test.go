package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
