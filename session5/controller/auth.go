package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	// get header request dari client
	authorization := c.GetHeader("Authorization")
	// coba kasih basic static auth di middleware
	if authorization !=  "12345" {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
		c.Abort()
	}

	// coba set variabel di middleware yang nantinya dapat diakses dari controller lain
	c.Set("user_name", "namaku")

	c.Next()
}
