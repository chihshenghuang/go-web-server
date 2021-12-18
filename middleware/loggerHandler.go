package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// logger middleware
func LoggerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		// before request

		ctx.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}
}