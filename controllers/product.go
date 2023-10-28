package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/models"
	"pharm-stock/utils/request"
	"pharm-stock/utils/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Interface beetween controller and routes
type ProductsControllerInterface interface {
	CreateProduct() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	SearchProduct() echo.HandlerFunc
}

// Connect into db and model
type ProductsController struct {
	config configs.Config
	model  models.ProductsModelInterface
}

// Create new instance from ProductsController
func NewProductsControllerInterface(m models.ProductsModelInterface) ProductsControllerInterface {
	return &ProductsController{
		model: m,
	}
}

// Create Product
func (pc *ProductsController) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.InsertProductRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Product input", errBind))
		}

		var newProduct = models.Products{}
		newProduct.Name = input.Name
		newProduct.Photo = input.Photo
		newProduct.IdCatProduct = input.IdCatProduct
		newProduct.MfDate = input.MfDate
		newProduct.ExpDate = input.ExpDate
		newProduct.BatchNumber = input.BatchNumber
		newProduct.UnitPrice = input.UnitPrice
		newProduct.Stock = input.Stock
		newProduct.Description = input.Description
		newProduct.IdDistributor = input.IdDistributor

		var res, errQuery = pc.model.Insert(newProduct)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery))
		}

		var insertResponse = response.ProductResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.Photo = res.Photo
		insertResponse.IdCatProduct = res.IdCatProduct
		insertResponse.MfDate = res.MfDate
		insertResponse.ExpDate = res.ExpDate
		insertResponse.BatchNumber = res.BatchNumber
		insertResponse.UnitPrice = res.UnitPrice
		insertResponse.Stock = res.Stock
		insertResponse.Description = res.Description
		insertResponse.IdDistributor = res.IdDistributor
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create Product", insertResponse))
	}
}

// Get All Product
func (pc *ProductsController) GetAllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = pc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all Product, ", err))
		}

		var getAllResponse []response.ProductResponse

		for _, product := range res {
			getAllResponse = append(getAllResponse, response.ProductResponse{
				Id:            product.Id,
				Name:          product.Name,
				Photo:         product.Photo,
				IdCatProduct:  product.IdCatProduct,
				MfDate:        product.MfDate,
				ExpDate:       product.ExpDate,
				BatchNumber:   product.BatchNumber,
				UnitPrice:     product.UnitPrice,
				Stock:         product.Stock,
				Description:   product.Description,
				IdDistributor: product.IdDistributor,
				CreatedAt:     product.CreatedAt,
				UpdatedAt:     product.UpdatedAt,
				DeletedAt:     product.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all Product, ", getAllResponse))
	}
}

// Update Product
func (pc *ProductsController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.Products{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Product input", errBind))
		}

		input.Id = paramId

		var res, errQuery = pc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery.Error()))
		}

		updateResponse := response.ProductResponse{}
		updateResponse.Id = res.Id
		updateResponse.Name = res.Name
		updateResponse.Photo = res.Photo
		updateResponse.IdCatProduct = res.IdCatProduct
		updateResponse.MfDate = res.MfDate
		updateResponse.ExpDate = res.ExpDate
		updateResponse.BatchNumber = res.BatchNumber
		updateResponse.UnitPrice = res.UnitPrice
		updateResponse.Stock = res.Stock
		updateResponse.Description = res.Description
		updateResponse.IdDistributor = res.IdDistributor
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}

// Delete Product
func (pc *ProductsController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		success, errQuery := pc.model.Delete(paramId)
		if !success {
			return c.JSON(http.StatusNotFound, response.FormatResponse("Product not found", errQuery))
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success delete Product", nil))
	}
}

// Searching
func (pc *ProductsController) SearchProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		products, err := pc.model.SearchProduct(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search category products, something happened", err))
		}

		var searchResponse []response.ProductResponse

		for _, product := range products {
			searchResponse = append(searchResponse, response.ProductResponse{
				Id:            product.Id,
				Name:          product.Name,
				Photo:         product.Photo,
				IdCatProduct:  product.IdCatProduct,
				MfDate:        product.MfDate,
				ExpDate:       product.ExpDate,
				BatchNumber:   product.BatchNumber,
				UnitPrice:     product.UnitPrice,
				Stock:         product.Stock,
				Description:   product.Description,
				IdDistributor: product.IdDistributor,
				CreatedAt:     product.CreatedAt,
				UpdatedAt:     product.UpdatedAt,
				DeletedAt:     product.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search category product success", searchResponse))
	}
}
