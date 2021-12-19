package main

import (
	"errors"
	"example/web-service-gin/middleware"
	"example/web-service-gin/model"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

// In memory cache for payloads data
var payloads = []model.Payload{
	{
		Title: "Valid App 1", 
		Version: "1.0.1", 
		Maintainers: []model.Maintainers{
			{
				Name: "firstmaintainer app1", 
				Email: "firstmaintainer@hotmail.com",
		  }, 
			{
				Name: "secondmaintainer app1", 
				Email: "secondmaintainer@gmail.com",
			},
		},
		Company: "Random Inc.",
		Website: "https://website.com",
		Source: "https://github.com/random/repo",
		License: "Apache-2.0",
		Description: `|
		### Interesting
		Title Some application content, and description`,
	},
	{
		Title: "Valid App 2", 
		Version: "0.0.1", 
		Maintainers: []model.Maintainers{
			{
				Name: "AppTwo Maintainer", 
				Email: "apptwo@hotmail.com",
		  }, 
		},
		Company: "Upbound Inc.",
		Website: "https://upbound.io",
		Source: "https://github.com/upbound/repo",
		License: "Apache-2.0",
		Description: `|
		### Why app 2 is the best
		Because it simply is...`,
	},
}

// getHealthStatus responds status 200 to show server is alive
func getHealthStatus(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{"data": "Service is healthy!"})
}

// getPayloads responds with the list of all payloads as JSON.
func getPayloads(ctx *gin.Context) {
	queryMap := getValidQueryParamsMap(ctx)
	results := getFilterPayloads(queryMap)

	ctx.YAML(http.StatusOK, results)
}

func getFilterPayloads(queryMap map[string][]string) []model.Payload {
	results := payloads

	// check each query key from query map 
	for queryKey, queryValueList := range queryMap {
		temp := []model.Payload{}
		// check each query value by query key
		for qvIdx := 0; qvIdx < len(queryValueList); qvIdx++ {
			// loop all payload data stored in memory
			for rIdx := 0; rIdx < len(results); rIdx++ {
				result := results[rIdx]
				
				// loop all struct field for each payload
				val := reflect.ValueOf(&result).Elem()
				for idx := 0; idx < val.NumField(); idx++ {
					// get field from struct
					field :=  strings.ToLower(val.Type().Field(idx).Name)
					if queryKey == field {
						fieldVal := val.Field(idx).Interface()
						if queryValueList[qvIdx] == fieldVal {
							temp = append(temp, result)
						}
					}
				}
			} 
		}
		results = temp
	}
	return results
}

func getValidQueryParamsMap(ctx *gin.Context) map[string][]string{	
	// get all query params from request
	queryMap := ctx.Request.URL.Query()

	validQueryMap := make(map[string][]string)
	
	payload := model.Payload{}
	val := reflect.ValueOf(&payload).Elem()

	// loop all property in payload struct to get valid query params
	for i := 0; i < val.NumField(); i++ {
		// get field from struct
		field :=  strings.ToLower(val.Type().Field(i).Name)

		if res, ok := queryMap[field]; ok {
			validQueryMap[field] = res
		}
	}

	return validQueryMap
}

// build error message from validation error
func getValidationError(verr validator.ValidationErrors) []ValidationError{
	errs := []ValidationError{}

	for _, ferr := range verr {
		err := ferr.ActualTag()
		if ferr.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, ferr.Param())
		}
		errs = append(errs, ValidationError{Field: ferr.Field(), Message:  validationErrorToText(ferr)})
	}

	return errs
}

func validationErrorToText(err validator.FieldError) string {
	switch err.ActualTag() {
	case "required":
		return fmt.Sprintf("Missing %s, %s is required", err.Field(), err.Field())
	case "email":
		return fmt.Sprintf("Invalid email format: %s", err.Value())
	}
	return fmt.Sprintf("%s is no valid", err.Field())
}

// postPayloads adds a payload from JSON received in the request body.
func postPayloads(ctx *gin.Context) {
	var newPayload model.Payload

	// Call ShouldBindYAML to bind the received YAML to
	// new payload.
	if err := ctx.ShouldBindYAML(&newPayload); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.YAML(http.StatusBadRequest, gin.H{"errors": getValidationError(verr)})
		  return
		}
		
		// Not a YAML format validation error
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	
	// Add the new payload to the slice.
	payloads = append(payloads, newPayload)
	ctx.YAML(http.StatusCreated, newPayload)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	
	// add logger middleware
	router.Use(middleware.LoggerHandler())

	// add authentication validation middleware
	router.Use(middleware.AuthenticationHandler())
	
	// add authorization validation middleware
	router.Use(middleware.AuthorizationHandler())

	// add rate limiter middleware
	router.Use(middleware.RateLimiterHandler())

	// api version 1
	v1 := router.Group("/v1")
	{
		v1.GET(("/health"), getHealthStatus)
		v1.GET("/payloads", getPayloads)
		v1.POST("/payloads", postPayloads)
	}

	// api version 2
	// to be continued in the future...

	return router
}

func main() {
	router := setupRouter()	
	router.Run("localhost:8080")
}

