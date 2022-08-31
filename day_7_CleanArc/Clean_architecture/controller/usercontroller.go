package usercontroller

import (
	"tesla/Entity"
	"tesla/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ControllerStudent struct {
	service1 service.ServiceStudent
}

func NewControllerStudent(service1 service.ServiceStudent) *ControllerStudent {
	return &ControllerStudent{service1}
}

func (co *ControllerStudent) GetStudentController(c echo.Context) error {
	users, err := co.service1.GetStudentService()

	if err != nil {
		log.Println("Get Student Controller Error", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (co *ControllerStudent) GetByIdController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	user, err := co.service1.GetStudentByIdService(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (co *ControllerStudent) CreateStudentController(c echo.Context) error {
	var newStudent entity.Student
	if err := c.Bind(&newStudent); err != nil {
		log.Println("Error Creating Student", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, errCreatedUser := co.service1.CreateStudentService(newStudent.Name, newStudent.Email)
	if errCreatedUser != nil {
		log.Println("Error Creating Student", errCreatedUser)
		return c.JSON(http.StatusBadRequest, errCreatedUser)
	}

	return c.JSON(http.StatusAccepted, createdUser)
}

func (co *ControllerStudent) UpdateStudentController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	var updateUser entity.Student
	if err := c.Bind(&updateUser); err != nil {
		log.Println("Error Updating Student", err)
		return c.JSON(http.StatusBadRequest, errId)
	}

	updatedUser, errUpdatedUser := co.service1.UpdateStudentService(id, updateUser.Name, updateUser.Email)
	if errUpdatedUser != nil {
		log.Println("Error Updating User", errUpdatedUser)
		return c.JSON(http.StatusBadRequest, errUpdatedUser)
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (co *ControllerStudent) DeleteStudentController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	_, err := co.service1.DeleteStudentService(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	response := "Delete Success"

	return c.JSON(http.StatusOK, response)
}