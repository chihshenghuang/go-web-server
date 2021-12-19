package main

import (
	"example/web-service-gin/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = setupRouter()
}

// test get health status api
func TestHealthRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/health", nil)
	router.ServeHTTP(w, req)
	
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
	router.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, payloads, got)
}

// test get payloads endpoint with title query string
func TestGetPayloadsQueryTitle(t *testing.T) {
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payloads?title=Valid App 1", nil)
	router.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, payloads[0:1], got)
}

// test get payloads endpoint with version and license query string
func TestGetPayloadsQueryVersionLicense(t *testing.T) {
	w :=  httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payloads?version=0.0.1&license=Apache-2.0", nil)
	router.ServeHTTP(w, req)

	// test status code
	assert.Equal(t, http.StatusOK, w.Code)

	// test content
	var got []model.Payload
	err := yaml.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, payloads[1:], got)
}


// test post payloads endpoint
// func TestPostPayloads(t *testing.T) {
// 	// read newPost yaml file
// 	newPostYaml, err := ioutil.ReadFile("./test/newPost.yaml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var newPost model.Payload
// 	err = yaml.Unmarshal(newPostYaml, &newPost)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
	
// 	newPostStr := fmt.Sprintf("%v", newPost)

// 	w :=  httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/v1/payloads", bytes.NewBufferString(newPostStr))
// 	req.Header.Add("Content-Type", "application/x-yaml;charset=utf-8")
// 	router.ServeHTTP(w, req)

// 	// test status code
// 	assert.Equal(t, http.StatusCreated, w.Code)

// 	// test content
// 	var got gin.H
// 	err = yaml.Unmarshal(w.Body.Bytes(), &got)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, newPost, got)
// }

