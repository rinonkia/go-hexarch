package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{},
		)
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
