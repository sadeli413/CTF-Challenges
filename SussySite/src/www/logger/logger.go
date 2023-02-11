package main

import (
	"log"
	"logger/app"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./assets/index.html")
	router.GET("/", app.Index)
	router.GET("/logs", app.Logs)
	if err := router.Run("localhost:7003"); err != nil {
		log.Fatal(err)
	}
}
