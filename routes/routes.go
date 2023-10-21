package routes

import (
	"pharm-stock/configs"
	"pharm-stock/controllers"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controllers.UserController, cfg configs.Config) {
	var user = e.Group("/users")

	user.POST("/login", uc.Login())
}
