package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHealthcheckHandler(t *testing.T) {
	mockResponse := `{"status":"ok"}`
	r := SetUpRouter()
	r.GET("/healthcheck", healthCheck)
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFlushKeys(t *testing.T) {
	mockResponse := `{"status":"flushed all keys"}`
	r := SetUpRouter()
	r.GET("/flush", flushKeys)
	req, _ := http.NewRequest("GET", "/flush", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateRecordHandler(t *testing.T) {
	mockResponse := `{"key":"mykey","status":"created","value":"ok"}`
	r := SetUpRouter()
	r.POST("/create", createRecord)
	jsonData := []byte(`{
		"key": "mykey",
		"value": "ok"
	}`)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteKeyHandler(t *testing.T) {
	mockResponse := `{"status":"deleted"}`
	r := SetUpRouter()
	r.DELETE("/delete", deleteRecord)
	req, _ := http.NewRequest("DELETE", "/delete?key=mykey", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
