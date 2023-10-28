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
type TransactionsControllerInterface interface {
	CreateTransaction() echo.HandlerFunc
	GetAllTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc
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

// Create Request Product
func (tc *TransactionsController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.InsertTransactionsRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Transaction input", errBind))
		}

		var newTransaction = models.Transactions{}
		newTransaction.IdEmployee = input.IdEmployee
		newTransaction.Type = input.Type

		var res, errQuery = tc.model.Insert(newTransaction)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery))
		}

		var insertResponse = response.TransactionResponse{}
		insertResponse.Id = res.Id
		insertResponse.IdEmployee = res.IdEmployee
		insertResponse.TotalQuantity = res.TotalQuantity
		insertResponse.TotalPrice = res.TotalPrice
		insertResponse.Type = res.Type
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create Request Product", insertResponse))
	}
}

// Get All Request Prodct
func (tc *TransactionsController) GetAllTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = tc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all Request Product, ", err))
		}

		var getAllResponse []response.TransactionResponse

		for _, Transaction := range res {
			getAllResponse = append(getAllResponse, response.TransactionResponse{
				Id:            Transaction.Id,
				IdEmployee:    Transaction.IdEmployee,
				TotalQuantity: Transaction.TotalQuantity,
				TotalPrice:    Transaction.TotalPrice,
				Type:          Transaction.Type,
				CreatedAt:     Transaction.CreatedAt,
				UpdatedAt:     Transaction.UpdatedAt,
				DeletedAt:     Transaction.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all Request Product, ", getAllResponse))
	}
}

// Update Request Product
func (tc *TransactionsController) UpdateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.Transactions{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Request Product input", errBind))
		}

		input.Id = paramId

		var res, errQuery = tc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery.Error()))
		}

		updateResponse := response.TransactionResponse{}
		updateResponse.Id = res.Id
		updateResponse.IdEmployee = res.IdEmployee
		updateResponse.TotalQuantity = res.TotalQuantity
		updateResponse.TotalPrice = res.TotalPrice
		updateResponse.Type = res.Type
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}

// Delete Request Product
func (tc *TransactionsController) DeleteTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		success, errQuery := tc.model.Delete(paramId)
		if !success {
			return c.JSON(http.StatusNotFound, response.FormatResponse("Request Product not found", errQuery))
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success delete Request Product", nil))
	}
}

// Searching
func (tc *TransactionsController) SearchTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		transactions, err := tc.model.SearchTransaction(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search category products, something happened", err))
		}

		var searchResponse []response.TransactionResponse

		for _, Transaction := range transactions {
			searchResponse = append(searchResponse, response.TransactionResponse{
				Id:            Transaction.Id,
				IdEmployee:    Transaction.IdEmployee,
				TotalQuantity: Transaction.TotalQuantity,
				TotalPrice:    Transaction.TotalPrice,
				Type:          Transaction.Type,
				CreatedAt:     Transaction.CreatedAt,
				UpdatedAt:     Transaction.UpdatedAt,
				DeletedAt:     Transaction.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search category product success", searchResponse))
	}
}
