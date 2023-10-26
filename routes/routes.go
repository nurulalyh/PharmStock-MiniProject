package routes

import (
	"pharm-stock/configs"
	"pharm-stock/controllers"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controllers.UsersControllerInterface, cfg configs.Config) {
	var user = e.Group("/users")

	user.POST("/admin", uc.CreateAdmin())
	user.POST("/login", uc.Login())

	// user.Use(authentication.JWTMiddleware)
	// user.Use(echojwt.JWT([]byte(cfg.Secret)))
	user.POST("", uc.CreateUser())
	user.GET("", uc.GetAllUsers())
	user.PUT("/:id", uc.UpdateUser())
	user.DELETE("/:id", uc.DeleteUser())
	user.GET("/search", uc.SearchUsers())
}

// func RouteCatProduct(e *echo.Echo, cpc controllers.CatProductControllerInterface, cfg configs.Config) {
// 	var user = e.Group("/catproducts")

// 	user.POST("", cpc.CreateCatProduct())
// 	user.GET("", cpc.GetAllCatProduct())
// 	user.GET("/:id", cpc.GetCatProductById())
// 	user.PUT("/:id", cpc.UpdateCatProduct())
// 	user.DELETE("/:id", cpc.DeleteCatProduct())
// }

// func RouteDistributor(e *echo.Echo, dc controllers.DistributorControllerInterface, cfg configs.Config) {
// 	var user = e.Group("/distributors")

// 	user.POST("", dc.CreateDistributor())
// 	user.GET("", dc.GetAllDistributor())
// 	user.GET("/:id", dc.GetDistributorById())
// 	user.PUT("/:id", dc.UpdateDistributor())
// 	user.DELETE("/:id", dc.DeleteDistributor())
// }

// func RouteReqProduct(e *echo.Echo, rpc controllers.ReqProductControllerInterface, cfg configs.Config) {
// 	var user = e.Group("/reqproducts")

// 	user.POST("", rpc.CreateReqProduct())
// 	user.GET("", rpc.GetAllReqProduct())
// 	user.GET("/:id", rpc.GetReqProductById())
// 	user.PUT("/:id", rpc.UpdateReqProduct())
// 	user.DELETE("/:id", rpc.DeleteReqProduct())
// }
