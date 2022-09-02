package main

import (
	"sample-app/controller"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middeleware"
	"net/http"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"massage":       "success to connect service",
		}
		return c.JSON(http.StatusOK, result)
	})

	cust := e.Group("customer")
	cust.GET("", controller.GetCustomer)
	cust.POST("", controller.CreateCustomer)
	cust.GET("/:id", controller.GetByID)
	cust.PUT("/:id", controller.UpdateCustomer)
	cust.DELETE("/:id", controller.DeleteCustomer)
	e.Logger.Fatal(e.Start("127.0.0.1:5002"))
	
}
