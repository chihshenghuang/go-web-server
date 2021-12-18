package main

import (
	"errors"
	"example/web-service-gin/middleware"
	"example/web-service-gin/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getHealthStatus responds status 200 to show server is alive
func getHealthStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Success")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// getAlbumbyID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

// build error message from validation error
func getValidationError(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = fmt.Sprintf("Missing %s, %s is %s!", f.StructField(), f.StructField(), err)
	}

	return errs
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(ctx *gin.Context) {
	var newAlbum model.Album

	// Call ShouldBindJSON to bind the received JSON to
	// new Albums.
	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": getValidationError(verr)})
		  return
		}

		// Not a validation error
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	
	// add logger middleware
	router.Use(middleware.LoggerHandler())

	// add authentication validation middleware
	router.Use(middleware.AuthenticationHandler())
	
	// add authorization validation middleware
	router.Use(middleware.AuthorizationHandler())

	// api version 1
	v1 := router.Group("/v1")
	{
		v1.GET(("/health"), getHealthStatus)
		v1.GET("/albums", getAlbums)
		v1.GET("/album/:id", getAlbumById)
		v1.POST("/albums", postAlbums)
	}

	// api version 2
	// to be continued in the future...

	return router
}

func main() {
	router := setupRouter()	
	router.Run("localhost:8080")
}

