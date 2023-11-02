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
type TransactionsControllerInterface interface {
	CreateTransaction() echo.HandlerFunc
	GetAllTransaction() echo.HandlerFunc
	SearchTransaction() echo.HandlerFunc
}

// Connect into db and model
type TransactionsController struct {
	config configs.Config
	model  models.TransactionsModelInterface
}

// Create new instance from TransactionsController
func NewTransactionsControllerInterface(m models.TransactionsModelInterface) TransactionsControllerInterface {
	return &TransactionsController{
		model: m,
	}
}

// Create Transaction
func (tc *TransactionsController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.TransactionsRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Transaction input", errBind.Error()))
		}

		var newTransaction = models.Transactions{}
		newTransaction.IdEmployee = input.IdEmployee
		newTransaction.Type = input.Type

		var res, errQuery = tc.model.Insert(newTransaction)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.InsertTransactionResponse{}
		insertResponse.Id = res.Id
		insertResponse.IdEmployee = res.IdEmployee
		insertResponse.TotalQuantity = res.TotalQuantity
		insertResponse.TotalPrice = res.TotalPrice
		insertResponse.Type = res.Type
		insertResponse.CreatedAt = res.CreatedAt

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create Transaction", insertResponse))
	}
}

// Get All Request Prodct
func (tc *TransactionsController) GetAllTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = tc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Transaction, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Transaction, ", res))
	}
}

// Searching
func (tc *TransactionsController) SearchTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		transactions, errQuery := tc.model.SearchTransaction(keyword, limit, offset)
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", errQuery.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", transactions))
	}
}
