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

		// TODO: use tools like Azure Application Insight to log the user request information
		// for each request in log middleware 
		ctx.Next()

		// log the latency
		latency := time.Since(t)
		log.Print(latency)

		// log the status
		status := ctx.Writer.Status()
		log.Println(status)
	}
}