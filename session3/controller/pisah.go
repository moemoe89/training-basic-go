package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Pisah(c *gin.Context) {
	c.String(http.StatusOK, "Hello pisah")
}
