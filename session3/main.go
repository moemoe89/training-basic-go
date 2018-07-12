package main

import (
	"net/http"
	"training-basic-go/session3/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	r.GET("/pisah", controller.Pisah)
	r.POST("/pisah-post", controller.Pisah)
	r.GET("/url-param", controller.UrlParam)
	r.GET("/url-path/:id", controller.UrlPath)
	r.POST("/input-post", controller.Inputpost)
	r.POST("/input-json", controller.InputJSON)

	r.Run(":8080")
}
