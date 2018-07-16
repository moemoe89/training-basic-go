package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"training-basic-go/session5/model"
)

func DeviceDetail(c *gin.Context) {

	id := c.Param("id")

	var device model.DeviceModel
	// contoh one to one join
	err := model.DB.Model(&device).Preload("User").Where("id = ?", id).First(&device).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data":    device,
	})
}

