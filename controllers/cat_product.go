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
type CatProductsControllerInterface interface {
	CreateCatProduct() echo.HandlerFunc
	GetAllCatProduct() echo.HandlerFunc
	UpdateCatProduct() echo.HandlerFunc
	DeleteCatProduct() echo.HandlerFunc
	SearchCatProduct() echo.HandlerFunc
}

// Connect into db and model
type CatProductsController struct {
	config configs.Config
	model  models.CatProductsModelInterface
}

// Create new instance from UserController
func NewCatProductsControllerInterface(m models.CatProductsModelInterface) CatProductsControllerInterface {
	return &CatProductsController{
		model: m,
	}
}

// Create Category Product
func (cpc *CatProductsController) CreateCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.InsertCategoryProductRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid category product input", errBind))
		}

		var newCatProduct = models.CatProducts{}
		newCatProduct.Name = input.Name

		var res, errQuery = cpc.model.Insert(newCatProduct)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.CatProductsResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create category product", insertResponse))
	}
}

// Get All Category Product
func (cpc *CatProductsController) GetAllCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = cpc.model.SelectAll(limit, offset)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all category product, ", err))
		}

		getAllResponse := response.CatProductsResponse{}
		for _, catProduct := range res {
			getAllResponse.Id = catProduct.Id
			getAllResponse.Name = catProduct.Name
			getAllResponse.CreatedAt = catProduct.CreatedAt
			getAllResponse.UpdatedAt = catProduct.UpdatedAt
			getAllResponse.DeletedAt = catProduct.DeletedAt
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all category product, ", getAllResponse))
	}
}

// Update Category Product
func (cpc *CatProductsController) UpdateCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.CatProducts{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid category product input", errBind))
		}

		input.Id = paramId

		var res, errQuery = cpc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery))
		}

		updateResponse := response.CatProductsResponse{}
		updateResponse.Id = res.Id
		updateResponse.Name = res.Name
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}

// Delete Category Product
func (cpc *CatProductsController) DeleteCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
	  var paramId = c.Param("id")

	  success, errQuery := cpc.model.Delete(paramId)
	  if !success {
		return c.JSON(http.StatusNotFound, response.FormatResponse("Category product not found", errQuery))
	  }

	  return c.JSON(http.StatusOK, response.FormatResponse("Success delete category product", nil))
	}
}

// Searching
func (cpc *CatProductsController) SearchCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		catProducts, err := cpc.model.SearchCatProduct(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search category products, something happened", err))
		}

		searchResponse := response.CatProductsResponse{}
		for _, catProduct := range catProducts{
			searchResponse.Id = catProduct.Id
			searchResponse.Name = catProduct.Name
			searchResponse.CreatedAt = catProduct.CreatedAt
			searchResponse.UpdatedAt = catProduct.UpdatedAt
			searchResponse.DeletedAt = catProduct.DeletedAt
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search category product success", searchResponse))
	}
}