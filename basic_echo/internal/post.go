package internal

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func PrintValue(c echo.Context) error {

	var response struct {
		ServerStatus string `json:"server_status"`
		Message      string `json:"message"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		EmpId        string `json:"emp_id"`
	}

	if err := c.Bind(&response); err != nil {
		response.ServerStatus = "Normal"
		response.Message = "request json error"
		return c.JSON(http.StatusOK, response)
	}

	response.ServerStatus = "Normal"
	response.Message = "Hello World"
	return c.JSON(http.StatusOK, response)
}

func Call3rdSystem(c echo.Context) error {

	var response struct {
		ServerStatus string      `json:"server_status"`
		Message      string      `json:"message"`
		Data         interface{} `json:"data"`
	}


	url := "https://jsonplaceholder.typicode.com/comments"
	id := c.QueryParam("id")

	//fmt.Println("id:>", id)

	if id != ""{
		url = fmt.Sprintf("%s/%s", url, id)
	}

	//fmt.Println("URL:>", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		response.ServerStatus = "Normal"
		response.Message = "error cannot init http new request"
		return c.JSON(http.StatusBadRequest, response)
	}

	req.Header.Set("Content-Type", "application/json")

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: t,
	}

	resp, err := client.Do(req)
	if err != nil {
		response.ServerStatus = "Normal"
		response.Message = "error cannot call 3rd system"
		return c.JSON(http.StatusBadRequest, response)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	response.ServerStatus = "Normal"
	//	response.Message = "error cannot read 3rd response"
	//	return c.JSON(http.StatusBadRequest, response)
	//}
	//fmt.Println("response Body:", string(body))

	type resp3rd struct {
		PostId int    `json:"postId"`
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}

	var resp3rdArr []resp3rd

	if err := json.NewDecoder(resp.Body).Decode(&resp3rdArr); err != nil {
		response.ServerStatus = "Normal"
		response.Message = "error cannot read 3rd response"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.ServerStatus = "Normal"
	response.Message = ""
	response.Data = resp3rdArr
	return c.JSON(http.StatusOK, response)
}
