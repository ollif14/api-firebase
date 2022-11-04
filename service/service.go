package service

import (
	"api-firebase/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
	"strings"
)

func EchoHTTPService() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// routes
	e.GET("/PHScale", GetDataPHScale())
	e.GET("/api/firebase-luis", GetDataPHScale())
	e.PUT("/api/update", UpdateFan())
	e.PUT("/api/update-fan/true", UpdateFanTrue())
	e.PUT("/api/update-fan/false", UpdateFanFalse())
	// e.POST("url", fucn())
	// e.DELETE("url", fucn())

	// run actual server
	e.Logger.Fatal(e.Start(":8999"))
}

func GetDataPHScale() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://project-aquascape-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/PHScale.json"
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.Result

		respponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(respponse, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateFan() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": true}"

		payload := strings.NewReader(body)

		url := "https://project-aquascape-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.Status

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateFanTrue() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": true}"

		payload := strings.NewReader(body)

		url := "https://project-aquascape-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.Status

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateFanFalse() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": false}"

		payload := strings.NewReader(body)

		url := "https://project-aquascape-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.Status

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}
