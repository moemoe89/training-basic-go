package main

import (
	"training-basic-go/session5/controller"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()
	r.POST("/user", controller.UserCreate)

	r.GET("/user", controller.UserList)
	r.GET("/user/:id", controller.UserDetail)
	r.GET("/device/:id", controller.DeviceDetail)

	// contoh grouping route
	grouping := r.Group("/grouping")
	// contoh implementasi middleware
	grouping.Use(controller.Auth)
	grouping.GET("/user", controller.UserListMiddleware)

	r.Run(":8080")

}

