package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/hello_worlds/service"
	"net/http"
)

type HelloWorldController struct {
	response response.Response
	service  service.HelloWorldService
}

func NewHelloWorldController(response response.Response, service service.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{
		response: response,
		service:  service,
	}
}

func (controller *HelloWorldController) Hello(c echo.Context) error {

	errorFlag := c.QueryParam("error")
	messageController := "hello from controller -"

	result, err := controller.service.Hello(c.Request().Context(), messageController, errorFlag)
	if err != nil {
		//TODO : create builder pattern
		return controller.
			response.
			Base(c, http.StatusInternalServerError, response.StatusFailed, result, err)
	}

	return controller.
		response.
		Base(c, http.StatusOK, response.StatusSuccess, result, err)
}
