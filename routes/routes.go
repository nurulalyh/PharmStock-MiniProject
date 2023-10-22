package routes

import (
	"pharm-stock/configs"
	"pharm-stock/controllers"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controllers.UserControllerInterface, cfg configs.Config) {
	var user = e.Group("/users")

	user.POST("", uc.CreateUser())
	user.POST("/login", uc.Login())
	user.GET("", uc.GetAllUsers())
	user.GET("/:id", uc.GetUserById())
	user.PUT("/:id", uc.UpdateUser())
}
