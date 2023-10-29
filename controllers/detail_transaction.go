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
type DetailTransactionsControllerInterface interface {
	CreateDetailTransaction() echo.HandlerFunc
	GetAllDetailTransaction() echo.HandlerFunc
	UpdateDetailTransaction() echo.HandlerFunc
	DeleteDetailTransaction() echo.HandlerFunc
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
		var input = request.InsertDetailTransactionsRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Detail Transaction input", errBind))
		}

		var newDetailTransaction = models.DetailTransactions{}
		newDetailTransaction.IdTransaction = input.IdTransaction
		newDetailTransaction.IdProduct = input.IdProduct
		newDetailTransaction.Quantity = input.Quantity
		newDetailTransaction.Price = 0
		
		var res, errQuery = dtc.model.Insert(newDetailTransaction)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery))
		}

		var insertResponse = response.DetailTransactionsResponse{}
		insertResponse.Id = res.Id
		insertResponse.IdTransaction = res.IdTransaction
		insertResponse.IdProduct = res.IdProduct
		insertResponse.Quantity = res.Quantity
		insertResponse.Price = res.Price
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create Detail Transaction", insertResponse))
	}
}

// Get All Detail Transaction
func (dtc *DetailTransactionsController) GetAllDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = dtc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all Detail Transaction, ", err))
		}

		var getAllResponse []response.DetailTransactionsResponse

		for _, DetailTransaction := range res {
			getAllResponse = append(getAllResponse, response.DetailTransactionsResponse{
				Id:          DetailTransaction.Id,
				IdTransaction:  DetailTransaction.IdTransaction,
				IdProduct: DetailTransaction.IdProduct,
				Quantity:    DetailTransaction.Quantity,
				Price:        DetailTransaction.Price,
				CreatedAt:   DetailTransaction.CreatedAt,
				UpdatedAt:   DetailTransaction.UpdatedAt,
				DeletedAt:   DetailTransaction.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all Detail Transaction, ", getAllResponse))
	}
}

// Update Detail Transaction
func (dtc *DetailTransactionsController) UpdateDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.DetailTransactions{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid Detail Transaction input", errBind))
		}

		input.Id = paramId

		var res, errQuery = dtc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery.Error()))
		}

		updateResponse := response.DetailTransactionsResponse{}
		updateResponse.Id = res.Id
		updateResponse.IdTransaction = res.IdTransaction
		updateResponse.IdProduct = res.IdProduct
		updateResponse.Quantity = res.Quantity
		updateResponse.Price = res.Price
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}

// Delete Detail Transaction
func (dtc *DetailTransactionsController) DeleteDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		success, errQuery := dtc.model.Delete(paramId)
		if !success {
			return c.JSON(http.StatusNotFound, response.FormatResponse("Detail Transaction not found", errQuery))
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success delete Detail Transaction", nil))
	}
}

// Searching
func (dtc *DetailTransactionsController) SearchDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		DetailTransactions, err := dtc.model.SearchDetailTransaction(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search category products, something happened", err))
		}

		var searchResponse []response.DetailTransactionsResponse

		for _, DetailTransaction := range DetailTransactions {
			searchResponse = append(searchResponse, response.DetailTransactionsResponse{
				Id:          DetailTransaction.Id,
				IdTransaction:  DetailTransaction.IdTransaction,
				IdProduct: DetailTransaction.IdProduct,
				Quantity:    DetailTransaction.Quantity,
				Price:        DetailTransaction.Price,
				CreatedAt:   DetailTransaction.CreatedAt,
				UpdatedAt:   DetailTransaction.UpdatedAt,
				DeletedAt:   DetailTransaction.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search category product success", searchResponse))
	}
}
