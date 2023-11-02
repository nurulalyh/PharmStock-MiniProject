package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/helper/authentication"
	"pharm-stock/models"
	"pharm-stock/utils/request"
	"pharm-stock/utils/response"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

// Create new instance from CatProductsController
func NewCatProductsControllerInterface(m models.CatProductsModelInterface) CatProductsControllerInterface {
	return &CatProductsController{
		model: m,
	}
}

// Create Category Product
func (cpc *CatProductsController) CreateCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, err := authentication.ExtractToken(userToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Invalid token", err.Error()))
			}

			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			var input = request.CategoryProductRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid category product input", errBind.Error()))
			}
	
			var newCatProduct = models.CatProducts{}
			newCatProduct.Name = input.Name
	
			var res, errQuery = cpc.model.Insert(newCatProduct)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
			}
	
			var insertResponse = response.InsertCatProductsResponse{}
			insertResponse.Id = res.Id
			insertResponse.Name = res.Name
			insertResponse.CreatedAt = res.CreatedAt
	
			return c.JSON(http.StatusCreated, helper.FormatResponse("success create category product", insertResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Get All Category Product
func (cpc *CatProductsController) GetAllCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = cpc.model.SelectAll(limit, offset)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all category product, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all category product, ", res))
	}
}

// Update Category Product
func (cpc *CatProductsController) UpdateCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, err := authentication.ExtractToken(userToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Invalid token", err.Error()))
			}

			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			var paramId = c.Param("id")
			var input = models.CatProducts{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid category product input", errBind.Error()))
			}
	
			input.Id = paramId
	
			var res, errQuery = cpc.model.Update(input)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
			}
	
			updateResponse := response.UpdateCatProductsResponse{}
			updateResponse.Id = res.Id
			updateResponse.Name = res.Name
			updateResponse.CreatedAt = res.CreatedAt
			updateResponse.UpdatedAt = res.UpdatedAt
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Delete Category Product
func (cpc *CatProductsController) DeleteCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, err := authentication.ExtractToken(userToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Invalid token", err.Error()))
			}

			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}
			
			var paramId = c.Param("id")
	
			success, errQuery := cpc.model.Delete(paramId)
			if !success {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Category product not found", errQuery.Error()))
			}
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success delete category product", nil))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Searching
func (cpc *CatProductsController) SearchCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		catProducts, errQuery := cpc.model.SearchCatProduct(keyword, limit, offset)
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", errQuery.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", catProducts))
	}
}
