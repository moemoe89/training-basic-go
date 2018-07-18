package main

import (
	"github.com/gin-gonic/gin"
	"training-basic-go/session6/controller"
)

func main(){

	r := gin.Default()
	r.GET("/user", controller.UserList)
	r.GET("/user/:id", controller.UserDetail)
	r.POST("/user", controller.UserCreate)
	r.PUT("/user/:id", controller.UserUpdate)
	r.DELETE("/user/:id", controller.UserDelete)
	r.POST("/user-form", controller.UserForm)

	r.Run(":8080")

}
