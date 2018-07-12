package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UrlPath(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "URL Path %s", id)
}
