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

func RouteCatProduct(e *echo.Echo, cpc controllers.CatProductsControllerInterface, cfg configs.Config) {
	var catProduct = e.Group("/catproducts")

	catProduct.POST("", cpc.CreateCatProduct())
	catProduct.GET("", cpc.GetAllCatProduct())
	catProduct.PUT("/:id", cpc.UpdateCatProduct())
	catProduct.DELETE("/:id", cpc.DeleteCatProduct())
	catProduct.GET("/search", cpc.SearchCatProduct())
}

func RouteDistributor(e *echo.Echo, dc controllers.DistributorsControllerInterface, cfg configs.Config) {
	var distributor = e.Group("/distributors")

	distributor.POST("", dc.CreateDistributor())
	distributor.GET("", dc.GetAllDistributor())
	distributor.PUT("/:id", dc.UpdateDistributor())
	distributor.DELETE("/:id", dc.DeleteDistributor())
	distributor.GET("/search", dc.SearchDistributor())
}

// func RouteReqProduct(e *echo.Echo, rpc controllers.ReqProductControllerInterface, cfg configs.Config) {
// 	var user = e.Group("/reqproducts")

// 	user.POST("", rpc.CreateReqProduct())
// 	user.GET("", rpc.GetAllReqProduct())
// 	user.GET("/:id", rpc.GetReqProductById())
// 	user.PUT("/:id", rpc.UpdateReqProduct())
// 	user.DELETE("/:id", rpc.DeleteReqProduct())
// }
