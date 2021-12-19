package main

import "example/web-service-gin/router"

func main() {
	router := router.SetupRouter()	
	router.Run("localhost:8080")
}

