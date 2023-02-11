package main

import (
	"log"
	"sussychat/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./assets/*.html")
	r.Static("/styles", "./styles/")
	r.Static("/images", "./images/")

	r.GET("/", app.GetHTML("index.html"))
	r.GET("/chat", app.GetHTML("chat.html"))
	r.GET("/frontend", app.Frontend)
	r.POST("/frontend", app.PostFrontend)
	r.GET("/uploads/:token", app.Uploads)

	if err := r.Run("localhost:7001"); err != nil {
		log.Fatal(err)
	}
}
