package main

import (
	"training-basic-go/session4/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/user", controller.User)
	r.GET("/user/:id", controller.UserDetail)
	r.POST("/user", controller.UserCreate)
	r.PUT("/user/:id", controller.UserUpdate)
	r.DELETE("/user/:id", controller.UserDelete)

	r.Run(":8080")

}
