package main

import (
	"fmt"
	"os"

	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)

	port := os.Getenv("PORT")

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}