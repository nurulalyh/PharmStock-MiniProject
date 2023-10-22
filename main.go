package main

import (
	"fmt"
	"pharm-stock/configs"
	"pharm-stock/controllers"
	"pharm-stock/models"
	"pharm-stock/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := models.InitModel(*config)
	models.Migrate(db)

	userModel := models.NewUsersModel(db)
	catProductModel := models.NewCatProductsModel(db)
	
	userController := controllers.NewUserControlInterface(userModel)
	catProductController := controllers.NewCatProductControllerInterface(catProductModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userController, *config)
	routes.RouteCatProduct(e, catProductController, *config)
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
