package controller

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"training-basic-go/session5/model"
)

func UserCreate(c *gin.Context) {

	var user model.UserModel
	err := c.ShouldBindJSON(&user)
	if err != nil {
		//if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	// contoh implementasi validasi
	err = user.Validation()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data": user,
	})
}


func UserList(c *gin.Context) {

	var user []model.UserModel
	// contoh one to many join
	err := model.DB.Preload("Devices").Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data": user,
	})
}

func UserDetail(c *gin.Context) {

	id := c.Param("id")

	var user model.UserModel
	// contoh one to many join
	err := model.DB.Where("id = ?", id).Preload("Devices").First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data":    user,
	})
}


func UserListMiddleware(c *gin.Context) {

	// mengambil variabel yang di set di middleware
	getUserData, ok := c.Get("user_name")
	if !ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "invalid data",
		})
		return
	}

	// assertion data ke bentuk semula (string)
	userData, ok := getUserData.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "invalid data",
		})
		return
	}

	// contoh cek isi variabel
	log.Print("nama: ", userData)

	var user []model.UserModel
	err := model.DB.Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data":    user,
	})
}
