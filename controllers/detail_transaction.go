package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/models"
	"pharm-stock/utils/request"
	"pharm-stock/utils/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Interface beetween controller and routes
type DetailTransactionsControllerInterface interface {
	CreateDetailTransaction() echo.HandlerFunc
	GetAllDetailTransaction() echo.HandlerFunc
	SearchDetailTransaction() echo.HandlerFunc
}

// Connect into db and model
type DetailTransactionsController struct {
	config configs.Config
	model  models.DetailTransactionsModelInterface
}

// Create new instance from DetailTransactionsController
func NewDetailTransactionsControllerInterface(m models.DetailTransactionsModelInterface) DetailTransactionsControllerInterface {
	return &DetailTransactionsController{
		model: m,
	}
}

// Create Detail Transaction
func (dtc *DetailTransactionsController) CreateDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.DetailTransactionsRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Detail Transaction input", errBind.Error()))
		}

		var newDetailTransaction = models.DetailTransactions{}
		newDetailTransaction.IdTransaction = input.IdTransaction
		newDetailTransaction.IdProduct = input.IdProduct
		newDetailTransaction.Quantity = input.Quantity

		var res, errQuery = dtc.model.Insert(newDetailTransaction)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.InsertDetailTransactionsResponse{}
		insertResponse.Id = res.Id
		insertResponse.IdTransaction = res.IdTransaction
		insertResponse.IdProduct = res.IdProduct
		insertResponse.Quantity = res.Quantity
		insertResponse.Price = res.Price
		insertResponse.CreatedAt = res.CreatedAt

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create Detail Transaction", insertResponse))
	}
}

// Get All Detail Transaction
func (dtc *DetailTransactionsController) GetAllDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = dtc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Detail Transaction, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Detail Transaction, ", res))
	}
}

// Searching
func (dtc *DetailTransactionsController) SearchDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		detailTransactions, errQuery := dtc.model.SearchDetailTransaction(keyword, limit, offset)
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", errQuery.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", detailTransactions))
	}
}
