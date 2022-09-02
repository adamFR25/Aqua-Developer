package controller

import (
	// "go/types"
	"log"
	"net/http"
	"strconv"

	// "fmt"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	ID    int `json:"id"`
	Name  string
	Email string
}

// var CustomerList = make(map[int]Customer)
var CustomerList = []Customer{
	{
		ID:    1,
		Name:  "Adam",
		Email: "adam@gmail.com",
	},
}

// CustomerList := []*Customer{}
var LastID int = 2

func GetCustomer(c echo.Context) (err error) {
	var result []Customer

	// DummyCustomer := Customer{
	// 	ID:    1,
	// 	Name:  "Adam",
	// 	Email: "colonel.travenn25@gmail.com",
	// }
	// result = append(result, DummyCustomer)
	for key, _ := range CustomerList {
		// tempCustomer := CustomerList(key)
		tempCustomer := CustomerList[key]
		result = append(result, tempCustomer)
	}
	log.Println("Get user ID", result[0].ID)
	return c.JSON(http.StatusOK, result)
}

func CreateCustomer(c echo.Context) error {
	req := new(Customer)
	log.Println("masuk ke create user")
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := Customer{
		ID:    LastID,
		Name:  req.Name,
		Email: req.Email,
	}

	CustomerList = append(CustomerList, user)
	log.Println("Create user ID")
	LastID++

	return c.JSON(http.StatusCreated, user)
}

func GetByID(c echo.Context) error {
	var result []Customer
	for key, _ := range CustomerList {
		// tempCustomer := CustomerList(key)
		tempCustomer := CustomerList[key]
		result = append(result, tempCustomer)
	}
	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	log.Println("Get user by ID", result[index-1].ID)

	return c.JSON(http.StatusAccepted, result[index-1])
}

func UpdateCustomer(c echo.Context) error {
	req := new(Customer)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	user := Customer{
		ID:    index,
		Name:  req.Name,
		Email: req.Email,
	}
	CustomerList[index-1] = user
	var result []Customer
	for key, _ := range CustomerList {
		// tempCustomer := CustomerList(key)
		tempCustomer := CustomerList[key]
		result = append(result, tempCustomer)
		// fmt.Println(result)
	}

	log.Println("Update user ID", result[index].ID)

	return c.JSON(http.StatusAccepted, user)

}

func DeleteCustomer(c echo.Context) error {
	req := new(Customer)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var result []Customer
	for key, _ := range CustomerList {
		// tempCustomer := CustomerList(key)
		tempCustomer := CustomerList[key]
		result = append(result, tempCustomer)
	}
	id := c.Param("id")
	index, _ := strconv.Atoi(id)

	// var indexSwap int
	var temp_res []Customer

	// temp_res = RemoveIndex(result,index)

	for i := 0; i < len(result); i++ {
		if i != index {
			temp_res = append(temp_res, result[i])
		}

	}

	return c.JSON(http.StatusAccepted, temp_res)

}

// func RemoveIndex(customer []Customer, index int) []Customer {
// 	return append(customer[:index], customer[index+1:]...)
// }

// melanjutkan endpoint customer get, get by id, update, dan delete.
// implementasikan log di dalam endpoint
// export postman coolection yg sudah dibuat
