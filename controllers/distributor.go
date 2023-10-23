package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/models"

	"github.com/labstack/echo/v4"
)

type DistributorControllerInterface interface {
	CreateDistributor() echo.HandlerFunc
	GetAllDistributor() echo.HandlerFunc
}

type DistributorController struct {
	config configs.Config
	model  models.DistributorModelInterface
}

func NewDistributorControllerInterface(m models.DistributorModelInterface) DistributorControllerInterface {
	return &DistributorController{
		model: m,
	}
}

func (dc *DistributorController) CreateDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = models.Distributor{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid distributor input", nil))
		}

		var res = dc.model.Insert(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create distributor", res))
	}
}

func (dc *DistributorController) GetAllDistributor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = dc.model.SelectAll()

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all distributor, ", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all distributor, ", res))
	}
}