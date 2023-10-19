package main

import (
	"fmt"
	"pharm-stock/configs"
	"pharm-stock/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := models.InitModel(*config)
	models.Migrate(db)

	// userModel := models.UsersModel{}
	// userModel.Init(db)
	// barangModel := model.NewBarangModel(db)

	// userControll := controller.UserController{}
	// userControll.InitUserController(userModel, *config)

	// barangControll := controller.NewBarangControllInterface(barangModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	// routes.RouteUser(e, userControll, *config)
	// routes.RouteBarang(e, barangControll, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
