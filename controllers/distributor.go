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
		var input = request.InsertDistributorsRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid distributor input", errBind))
		}

		var newDistributor = models.Distributors{}
		newDistributor.Name = input.Name

		var res, errQuery = dc.model.Insert(newDistributor)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.DistributorResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create distributor", insertResponse))
	}
}

// Get all distributor
func (dc *DistributorsController) GetAllDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = dc.model.SelectAll(limit, offset)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all distributor, ", err))
		}

		var getAllResponse []response.DistributorResponse 

		for _, distributor := range res {
			getAllResponse = append(getAllResponse, response.DistributorResponse{
				Id:        distributor.Id,
				Name:      distributor.Name,
				CreatedAt: distributor.CreatedAt,
				UpdatedAt: distributor.UpdatedAt,
				DeletedAt: distributor.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all distributor, ", getAllResponse))
	}
}

//Update distributor
func (dc *DistributorsController) UpdateDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.Distributors{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid distributor input", errBind))
		}

		input.Id = paramId

		var res, errQuery = dc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery))
		}

		updateResponse := response.DistributorResponse{}
		updateResponse.Id = res.Id
		updateResponse.Name = res.Name
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}


// Delete Distributor
func (dc *DistributorsController) DeleteDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
	  var paramId = c.Param("id")

	  success, errQuery := dc.model.Delete(paramId)
	  if !success {
		return c.JSON(http.StatusNotFound, response.FormatResponse("distributor not found", errQuery))
	  }

	  return c.JSON(http.StatusOK, response.FormatResponse("Success delete distributor", nil))
	}
}

// Searching
func (dc *DistributorsController) SearchDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		distributors, err := dc.model.SearchDistributor(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search category products, something happened", err))
		}

		var searchResponse []response.DistributorResponse

		for _, distributor := range distributors {
			searchResponse = append(searchResponse, response.DistributorResponse{
				Id:        distributor.Id,
				Name:      distributor.Name,
				CreatedAt: distributor.CreatedAt,
				UpdatedAt: distributor.UpdatedAt,
				DeletedAt: distributor.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search category product success", searchResponse))
	}
}