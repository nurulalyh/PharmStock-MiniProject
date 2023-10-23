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
	user.GET("", cpc.GetAllCatProduct())
	user.GET("/:id", cpc.GetCatProductById())
	user.PUT("/:id", cpc.UpdateCatProduct())
	user.DELETE("/:id", cpc.DeleteCatProduct())
}

func RouteDistributor(e *echo.Echo, dc controllers.DistributorControllerInterface, cfg configs.Config) {
	var user = e.Group("/distributor")

	user.POST("", dc.CreateDistributor())
	user.GET("", dc.GetAllDistributor())
	user.GET("/:id", dc.GetDistributorById())
	user.PUT("/:id", dc.UpdateDistributor())
	user.DELETE("/:id", dc.DeleteDistributor())
}
