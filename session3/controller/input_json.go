package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Input interface{} `json:"input"`
	Ouput struct {
		InsideOutput string `json:"inside_output"`
	} `json:"output"`
}

func InputJSON(c *gin.Context) {

	// binding input body (json) ke struct
	var user user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusOK, "Invalid")
		return
	}

	var resp ResponseStruct
	resp.Input = user
	resp.Ouput.InsideOutput = "isinya"

	c.JSON(http.StatusOK, resp)
}
