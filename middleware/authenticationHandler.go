package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authentication validation:
		// if authentication is failed, return 401 status response immediately
		if isAuthenticated(ctx) {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, "Authentication failed")
	  	ctx.Abort()
		}
	}
}

func isAuthenticated(ctx *gin.Context) bool {
	// check user's json token is valid
	// if use session id, check session id is valid

	// mockAuthNValidation()
	// assume user pass the authN validation
	return true;
}

