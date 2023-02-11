package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHTML(file string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, file, nil)
	}
}

func tok2H(token string) gin.H {
	if len(token) == tokenSize {
		return gin.H{
			"image": "/uploads/" + token,
		}
	}
	return nil
}

func Frontend(c *gin.Context) {
	obj := tok2H("")
	c.HTML(http.StatusOK, "frontend.html", obj)
}

func PostFrontend(c *gin.Context) {
	body := wget(c.PostForm("sussyAPI"))
	if len(body) == 0 {
		c.Redirect(http.StatusFound, "/frontend")
		return
	}

	token, err := save(body)
	if err != nil {
		c.Redirect(http.StatusFound, "/frontend")
		return
	}

	obj := tok2H(token)
	c.HTML(http.StatusOK, "frontend.html", obj)
}

func Uploads(c *gin.Context) {
	token := c.Param("token")
	body := load(token)
	c.Data(http.StatusOK, "", body)
}
