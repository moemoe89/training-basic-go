package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UrlParam(c *gin.Context) {

	name := c.Query("name")
	address := c.Query("address")

	c.String(http.StatusOK, "Hello pisah %s %s", name, address)
}
