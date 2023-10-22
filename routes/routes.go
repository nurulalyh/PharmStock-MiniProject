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
	user.DELETE("/:id", uc.DeleteUser())
}

func RouteCatProduct(e *echo.Echo, cpc controllers.CatProductControllerInterface, cfg configs.Config) {
	var user = e.Group("/catproducts")

	user.POST("", cpc.CreateCatProduct())
	// user.POST("/login", cpc.Login())
	// user.GET("", cpc.GetAllUsers())
	// user.GET("/:id", cpc.GetUserById())
	// user.PUT("/:id", cpc.UpdateUser())
	// user.DELETE("/:id", cpc.DeleteUser())
}
