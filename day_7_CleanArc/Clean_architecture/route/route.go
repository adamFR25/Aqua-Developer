package route

import (
	"tesla/repository"
	"tesla/service"
	"tesla/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *echo.Echo {
	//Initialize Echo
	e := echo.New()

	//Student
	StudentRepository := repository.NewRepositoryUser(db)
	StudentService := service.NewUserService(*StudentRepository)
	StudentController := usercontroller.NewControllerStudent(*StudentService)

	e.GET("/users", StudentController.GetStudentController)
	e.GET("/user/:id", StudentController.GetByIdController)
	e.POST("/newuser", StudentController.CreateStudentController)
	e.PUT("/updateuser/:id", StudentController.UpdateStudentController)
	e.DELETE("/deleteuser/:id", StudentController.DeleteStudentController)

	return e
}
