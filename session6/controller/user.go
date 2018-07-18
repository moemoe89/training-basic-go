package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"training6/model"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {

	name := c.Query("name")
	address := c.Query("address")

	// setup request method (get/post/put/delete/etc), url tujuan, dan body (karena get maka tidak ada body)
	req, err := http.NewRequest("GET", "https://peripi.herokuapp.com/user", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// contoh set header
	req.Header.Set("Authorization", "12345")

	// contoh set url parameter (query) ex: ?name=xxx&address=yyy
	queryValues := req.URL.Query()
	queryValues.Set("name", name)
	queryValues.Set("address", address)
	req.URL.RawQuery = queryValues.Encode()

	// set variabel http client untuk mengeksekusi request
	httpClient := &http.Client{}

	// eksekusi request
	resp, err := httpClient.Do(req)
	// jangan lupa selalu close body response
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	// contoh get http status code response (200/201/400/500/etc)
	var statusCode = resp.StatusCode

	var response model.UserListModel

	// assign body value ke struct yang kita definisikan sesuai response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}

func UserDetail(c *gin.Context) {

	id := c.Param("id")

	req, err := http.NewRequest("GET", "https://peripi.herokuapp.com/user/"+id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	req.Header.Set("Authorization", "12345")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var statusCode = resp.StatusCode

	var response model.UserDetailModel

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}

func UserCreate(c *gin.Context) {

	var userForm model.UserForm
	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	data, err := json.Marshal(userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	body := bytes.NewBuffer(data)
	defer body.Reset()

	req, err := http.NewRequest("POST", "https://peripi.herokuapp.com/user", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	req.Header.Set("Authorization", "12345")
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var statusCode = resp.StatusCode

	var response model.UserCreateModel

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}

func UserUpdate(c *gin.Context) {

	id := c.Param("id")

	var userForm model.UserForm
	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	data, err := json.Marshal(userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	body := bytes.NewBuffer(data)
	defer body.Reset()

	req, err := http.NewRequest("PUT", "https://peripi.herokuapp.com/user/"+id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	req.Header.Set("Authorization", "12345")
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var statusCode = resp.StatusCode

	var response model.UserCreateModel

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}

func UserDelete(c *gin.Context) {

	id := c.Param("id")

	req, err := http.NewRequest("DELETE", "https://peripi.herokuapp.com/user/"+id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	req.Header.Set("Authorization", "12345")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var statusCode = resp.StatusCode

	var response model.UserDeleteModel

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}

func UserForm(c *gin.Context) {

	var userForm model.UserForm
	err := c.ShouldBind(&userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	form := url.Values{}
	form.Add("name", userForm.Name)
	form.Add("address", userForm.Address)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", "https://peripi.herokuapp.com/user-form", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	req.Header.Set("Authorization", "12345")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	var statusCode = resp.StatusCode

	var response model.UserCreateModel

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, response)
}