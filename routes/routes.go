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

func RouteReqProduct(e *echo.Echo, rpc controllers.ReqProductsControllerInterface, cfg configs.Config) {
	var reqProduct = e.Group("/reqproducts")

	reqProduct.POST("", rpc.CreateReqProduct())
	reqProduct.GET("", rpc.GetAllReqProduct())
	reqProduct.PUT("/:id", rpc.UpdateReqProduct())
	reqProduct.GET("/search", rpc.SearchReqProduct())
}

func RouteTransaction(e *echo.Echo, tc controllers.TransactionsControllerInterface, cfg configs.Config) {
	var transaction = e.Group("/transactions")

	transaction.POST("", tc.CreateTransaction())
	transaction.GET("", tc.GetAllTransaction())
	transaction.GET("/search", tc.SearchTransaction())
}

func RouteProduct(e *echo.Echo, pc controllers.ProductsControllerInterface, cfg configs.Config) {
	var product = e.Group("/products")

	product.POST("", pc.CreateProduct())
	product.GET("", pc.GetAllProduct())
	product.PUT("/:id", pc.UpdateProduct())
	product.DELETE("/:id", pc.DeleteProduct())
	product.GET("/search", pc.SearchProduct())
}

func RouteDetailTransaction(e *echo.Echo, dtc controllers.DetailTransactionsControllerInterface, cfg configs.Config) {
	var detailTransaction = e.Group("/detailTransactions")

	detailTransaction.POST("", dtc.CreateDetailTransaction())
	detailTransaction.GET("", dtc.GetAllDetailTransaction())
	detailTransaction.PUT("/:id", dtc.UpdateDetailTransaction())
	detailTransaction.DELETE("/:id", dtc.DeleteDetailTransaction())
	detailTransaction.GET("/search", dtc.SearchDetailTransaction())
}
