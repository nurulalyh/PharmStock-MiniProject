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
	distributorModel := models.NewDistributorsModel(db)
	reqProductModel := models.NewReqProductsModel(db)
	transactionModel := models.NewTransactionsModel(db)
	productModel := models.NewProductsModel(db)
	detailTransactionModel := models.NewDetailTransactionsModel(db)

	userController := controllers.NewUsersControlInterface(userModel)
	catProductController := controllers.NewCatProductsControllerInterface(catProductModel)
	distributorController := controllers.NewDistributorControllerInterface(distributorModel)
	reqProductController := controllers.NewReqProductsControllerInterface(reqProductModel)
	transactionController := controllers.NewTransactionsControllerInterface(transactionModel)
	productController := controllers.NewProductsControllerInterface(productModel)
	detailTransactionController := controllers.NewDetailTransactionsControllerInterface(detailTransactionModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userController, *config)
	routes.RouteCatProduct(e, catProductController, *config)
	routes.RouteDistributor(e, distributorController, *config)
	routes.RouteReqProduct(e, reqProductController, *config)
	routes.RouteTransaction(e, transactionController, *config)
	routes.RouteProduct(e, productController, *config)
	routes.RouteDetailTransaction(e, detailTransactionController, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
