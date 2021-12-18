package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authorization validation:
		// if authorization is failed, return 403 status response immediately
		if isAuthorized(ctx) {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, "Authorization failed")
	  	ctx.Abort()
		}
	}
}

func isAuthorized(ctx *gin.Context) bool {
	// check if user has permission to access the resource

	// mockAuthZValidation()
	// assume user pass the authZ validation
	return true;
}