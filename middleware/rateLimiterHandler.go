package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RateLimiterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isPermitForRequest(ctx) {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusTooManyRequests, "Too many request")
	  	ctx.Abort()
		}
	}
}

func isPermitForRequest(ctx *gin.Context) bool {
	// use method like token bucket to check if the request is allowed

	// assuem the request is allowed
	return true
}