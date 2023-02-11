package main

import (
	"crypto/rand"
	"log"
  "math/big"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
  router.GET("/api", api)
  if err := router.Run("localhost:7002"); err != nil {
    log.Fatal(err)
  }
}

func api(c *gin.Context) {
  if filename, err := randFile(); err != nil {
    c.File("images/neco-arc00.png")
  } else {
    c.File(filename)
  }
}

func randFile() (string, error) {
	files := []string{
    "images/neco-arc00.png",
		"images/neco-arc01.png",
		"images/neco-arc02.png",
		"images/neco-arc00.jpg",
		"images/neco-arc01.jpg",
		"images/neco-arc02.jpg",
		"images/neco-arc03.jpg",
		"images/neco-arc04.jpg",
		"images/neco-arc05.jpg",
		"images/neco-arc06.jpg",
		"images/neco-arc00.gif",
		"images/neco-arc01.gif",
		"images/neco-arc02.gif",
		"images/neco-arc03.gif",
		"images/neco-arc04.gif",
		"images/neco-arc05.gif",
		"images/neco-arc06.gif",
		"images/neco-arc07.gif",
		"images/neco-arc08.gif",
	}

  i, err := rand.Int(rand.Reader, big.NewInt(int64(len(files))))
  if err != nil {
    return "", err
  }

	return files[i.Int64()], nil
}
