package controller

import (
	"github.com/gin-gonic/gin"
	"training4/model"
	"net/http"
	"strconv"
)

func User(c *gin.Context) {

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
		"data": user,
	})
}

func UserDetail(c *gin.Context) {

	id := c.Param("id")

	var user model.UserModel
	err := model.DB.Where("id = ?", id).First(&user).Error
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


func UserUpdate(c *gin.Context) {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var user model.UserModel
	err = c.ShouldBindJSON(&user)
	if err != nil {
		//if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	user.ID = id

	err = model.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data": user,
	})
}

func UserDelete(c *gin.Context) {

	id := c.Param("id")

	var user model.UserModel
	err := model.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}