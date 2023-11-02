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
type ReqProductsControllerInterface interface {
	CreateReqProduct() echo.HandlerFunc
	GetAllReqProduct() echo.HandlerFunc
	UpdateReqProduct() echo.HandlerFunc
	SearchReqProduct() echo.HandlerFunc
}

// Connect into db and model
type ReqProductsController struct {
	config configs.Config
	model  models.ReqProductsModelInterface
}

// Create new instance from ReqProductsController
func NewReqProductsControllerInterface(m models.ReqProductsModelInterface) ReqProductsControllerInterface {
	return &ReqProductsController{
		model: m,
	}
}

// Create Request Product
func (rpc *ReqProductsController) CreateReqProduct() echo.HandlerFunc {
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

			if role != "apoteker" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			var input = request.InsertReqProductRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", errBind.Error()))
			}
	
			var newReqProduct = models.ReqProducts{}
			newReqProduct.IdEmployee = input.IdEmployee
			newReqProduct.ProductName = input.ProductName
			newReqProduct.Quantity = input.Quantity
			newReqProduct.Note = input.Note
	
			var res, errQuery = rpc.model.Insert(newReqProduct)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
			}
	
			var insertResponse = response.InsertReqProductResponse{}
			insertResponse.Id = res.Id
			insertResponse.IdEmployee = res.IdEmployee
			insertResponse.ProductName = res.ProductName
			insertResponse.Quantity = res.Quantity
			insertResponse.Note = res.Note
			insertResponse.StatusReq = res.StatusReq
			insertResponse.CreatedAt = res.CreatedAt
	
			return c.JSON(http.StatusCreated, helper.FormatResponse("success create Request Product", insertResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Get All Request Prodct
func (rpc *ReqProductsController) GetAllReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = rpc.model.SelectAll(limit, offset)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Request Product, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Request Product, ", res))
	}
}

// Update Request Product
func (rpc *ReqProductsController) UpdateReqProduct() echo.HandlerFunc {
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
			var input = request.UpdateReqProductRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", errBind.Error()))
			}
	
			input.Id = paramId
	
			update := models.ReqProducts{}
			update.Id = input.Id
			update.StatusReq = input.StatusReq
	
			var res, errQuery = rpc.model.Update(update)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
			}
	
			updateResponse := response.UpdateReqProductResponse{}
			updateResponse.Id = res.Id
			updateResponse.IdEmployee = res.IdEmployee
			updateResponse.ProductName = res.ProductName
			updateResponse.Quantity = res.Quantity
			updateResponse.Note = res.Note
			updateResponse.StatusReq = res.StatusReq
			updateResponse.CreatedAt = res.CreatedAt
			updateResponse.UpdatedAt = res.UpdatedAt
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Searching
func (rpc *ReqProductsController) SearchReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		reqProducts, errQuery := rpc.model.SearchReqProduct(keyword, limit, offset)
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", errQuery.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", reqProducts))
	}
}
