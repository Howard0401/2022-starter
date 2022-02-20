package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"msg": "ok"},
		)
	})
	r.Run()
}
