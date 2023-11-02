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
type DistributorsControllerInterface interface {
	CreateDistributor() echo.HandlerFunc
	GetAllDistributor() echo.HandlerFunc
	UpdateDistributor() echo.HandlerFunc
	DeleteDistributor() echo.HandlerFunc
	SearchDistributor() echo.HandlerFunc
}

// Connect into db and model
type DistributorsController struct {
	config configs.Config
	model  models.DistributorModelInterface
}

// Create new instance from DistributorsController
func NewDistributorControllerInterface(m models.DistributorModelInterface) DistributorsControllerInterface {
	return &DistributorsController{
		model: m,
	}
}

// Create Distributor
func (dc *DistributorsController) CreateDistributor() echo.HandlerFunc {
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

			var input = request.DistributorsRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid distributor input", errBind.Error()))
			}
	
			var newDistributor = models.Distributors{}
			newDistributor.Name = input.Name
	
			var res, errQuery = dc.model.Insert(newDistributor)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
			}
	
			var insertResponse = response.InsertDistributorResponse{}
			insertResponse.Id = res.Id
			insertResponse.Name = res.Name
			insertResponse.CreatedAt = res.CreatedAt
	
			return c.JSON(http.StatusCreated, helper.FormatResponse("success create distributor", insertResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Get all distributor
func (dc *DistributorsController) GetAllDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = dc.model.SelectAll(limit, offset)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all distributor, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all distributor, ", res))
	}
}

// Update distributor
func (dc *DistributorsController) UpdateDistributor() echo.HandlerFunc {
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
			var input = models.Distributors{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid distributor input", errBind.Error()))
			}
	
			input.Id = paramId
	
			var res, errQuery = dc.model.Update(input)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
			}
	
			updateResponse := response.UpdateDistributorResponse{}
			updateResponse.Id = res.Id
			updateResponse.Name = res.Name
			updateResponse.CreatedAt = res.CreatedAt
			updateResponse.UpdatedAt = res.UpdatedAt
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))

	}
}

// Delete Distributor
func (dc *DistributorsController) DeleteDistributor() echo.HandlerFunc {
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
	
			success, errQuery := dc.model.Delete(paramId)
			if !success {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("distributor not found", errQuery.Error()))
			}
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success delete distributor", nil))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Searching
func (dc *DistributorsController) SearchDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		distributors, errQuery := dc.model.SearchDistributor(keyword, limit, offset)
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", errQuery.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", distributors))
	}
}
