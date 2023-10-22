package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/models"

	"github.com/labstack/echo/v4"
)

type CatProductControllerInterface interface {
	CreateCatProduct() echo.HandlerFunc
	GetAllCatProduct() echo.HandlerFunc
}

type CatProductController struct {
	config configs.Config
	model  models.CatProductModelInterface
}

func NewCatProductControllerInterface(m models.CatProductModelInterface) CatProductControllerInterface {
	return &CatProductController{
		model: m,
	}
}

func (cpc *CatProductController) CreateCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = models.CatProduct{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid category product input", nil))
		}

		var res = cpc.model.Insert(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create category product", res))
	}
}

func (cpc *CatProductController) GetAllCatProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = cpc.model.SelectAll()

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all category product, ", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all category product, ", res))
	}
}