package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func Inputpost(c *gin.Context) {

	// kita bisa get data form-post / url encoded dengan menggunakan binding ke struct
	var user user
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusOK, "Invalid")
		return
	}

	name    := user.Name
	address := user.Address

	//name := c.PostForm("name")
	//address := c.PostForm("address")

	c.String(http.StatusOK, "Input body %s %s", name, address)
}
