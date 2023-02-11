package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Logs(c *gin.Context) {
	file := "./logs/" + c.Query("file")
	if _, err := os.Stat(file); err != nil {
		c.Data(http.StatusOK, "application/x-www-form-urlencoded", []byte{})
	} else {
		c.File(file)
	}
}
