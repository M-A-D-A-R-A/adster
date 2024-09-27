package handler

import (
	"microservices/src/input"
	"microservices/src/service"
	"net/http"
	// "fmt"

	"github.com/labstack/echo/v4"
)

type targetHandler struct {
	service service.TargetService
}

func NewTargetHandler(service service.TargetService) *targetHandler {
	return &targetHandler{service}
}

func (h *targetHandler) GetAllFiles(c echo.Context) error {

		// Read the request body into a ForecastData struct
		var forecastRequest input.ForecastRequest
		if err := c.Bind(&forecastRequest); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request payload")
		}
	
		// Send a request to the third-party API
		responseData, err := h.service.GetForcast(forecastRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to send request to third-party API")
		}
	
		return c.JSON(http.StatusOK, responseData) // Return response data as JSON

    // files, err := h.service.GetForcast()

    // if err != nil {
    //     return c.String(http.StatusBadRequest, "Failed to get files data")
    // }
    // return c.JSON(http.StatusOK, files) // Return files as JSON
}


func (h *targetHandler) GetForcast(c echo.Context) error{
	var forecastRequest input.ForecastRequest
	files, err := h.service.GetForcast(forecastRequest)
    if err != nil {
        return c.String(http.StatusBadRequest, "Failed to get files data")
    }
    return c.JSON(http.StatusOK, files) // Return files as JSON
}



// func (h *targetHandler) sendForecastRequest(c echo.Context) error {

// 	// Read the request body into a ForecastData struct
// 	var forecastRequest input.ForecastRequest
// 	if err := c.Bind(&forecastRequest); err != nil {
// 		return c.String(http.StatusBadRequest, "Invalid request payload")
// 	}

// 	// Send a request to the third-party API
// 	// responseData, err := h.service.sendForecastRequest(&forecastRequest)
// 	responseData, err = h.service.GetForcast()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "Failed to send request to third-party API")
// 	}

// 	return c.JSON(http.StatusOK, responseData) // Return response data as JSON

// // files, err := h.service.GetForcast()

// // if err != nil {
// //     return c.String(http.StatusBadRequest, "Failed to get files data")
// // }
// // return c.JSON(http.StatusOK, files) // Return files as JSON
// }