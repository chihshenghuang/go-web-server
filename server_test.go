package main

import (
	"bytes"
	"example/web-service-gin/model"
	"example/web-service-gin/router"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var testRouter *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	testRouter = router.SetupRouter()
}

// test get health status api
func TestHealthRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/health", nil)
	testRouter.ServeHTTP(w, req)
	
	// test status code
	assert.Equal(t, http.StatusOK, w.Code)
	
	// test content
	var got gin.H 
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, gin.H{"data":"Service is healthy!"}, got)
}

// test get payloads endpoint
func TestGetPayloads(t *testing.T) {
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payloads", nil)
	testRouter.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, router.Payloads, got)
}

// test get payloads endpoint with title query string
func TestGetPayloadsQueryTitle(t *testing.T) {
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payloads?title=Valid App 1", nil)
	testRouter.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, router.Payloads[0:1], got)
}

// test get payloads endpoint with version and license query string
func TestGetPayloadsQueryVersionLicense(t *testing.T) {
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payloads?version=0.0.1&license=Apache-2.0", nil)
	testRouter.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, router.Payloads[1:2], got)
}

// test post payloads endpoint
func TestPostPayloads(t *testing.T) {
	// read newPost yaml file
	newPostYaml, err := ioutil.ReadFile("./test/newPost.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var newPost model.Payload
	err = yaml.Unmarshal(newPostYaml, &newPost)
	if err != nil {
		t.Fatal(err)
	}
	
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/payloads", bytes.NewBuffer(newPostYaml))
	req.Header.Add("Content-Type", "application/x-yaml;charset=utf-8")
	testRouter.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// test content
	var got model.Payload
	err = yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, newPost, got)
}