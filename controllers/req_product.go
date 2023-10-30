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
type ReqProductsControllerInterface interface {
	CreateReqProduct() echo.HandlerFunc
	GetAllReqProduct() echo.HandlerFunc
	UpdateReqProduct() echo.HandlerFunc
	DeleteReqProduct() echo.HandlerFunc
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
		var input = request.InsertReqProductRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", errBind))
		}

		var newReqProduct = models.ReqProducts{}
		newReqProduct.IdEmployee = input.IdEmployee
		newReqProduct.ProductName = input.ProductName
		newReqProduct.Quantity = input.Quantity
		newReqProduct.Note = input.Note
		newReqProduct.StatusReq = "processed"

		var res, errQuery = rpc.model.Insert(newReqProduct)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery))
		}

		var insertResponse = response.ReqProductResponse{}
		insertResponse.Id = res.Id
		insertResponse.IdEmployee = res.IdEmployee
		insertResponse.ProductName = res.ProductName
		insertResponse.Quantity = res.Quantity
		insertResponse.Note = res.Note
		insertResponse.StatusReq = res.StatusReq
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create Request Product", insertResponse))
	}
}

// Get All Request Prodct
func (rpc *ReqProductsController) GetAllReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = rpc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Request Product, ", err))
		}

		var getAllResponse []response.ReqProductResponse

		for _, reqProduct := range res {
			getAllResponse = append(getAllResponse, response.ReqProductResponse{
				Id:          reqProduct.Id,
				IdEmployee:  reqProduct.IdEmployee,
				ProductName: reqProduct.ProductName,
				Quantity:    reqProduct.Quantity,
				Note:        reqProduct.Note,
				StatusReq:   reqProduct.StatusReq,
				CreatedAt:   reqProduct.CreatedAt,
				UpdatedAt:   reqProduct.UpdatedAt,
				DeletedAt:   reqProduct.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Request Product, ", getAllResponse))
	}
}

// Update Request Product
func (rpc *ReqProductsController) UpdateReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.ReqProducts{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", errBind))
		}

		input.Id = paramId

		var res, errQuery = rpc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
		}

		updateResponse := response.ReqProductResponse{}
		updateResponse.Id = res.Id
		updateResponse.IdEmployee = res.IdEmployee
		updateResponse.ProductName = res.ProductName
		updateResponse.Quantity = res.Quantity
		updateResponse.Note = res.Note
		updateResponse.StatusReq = res.StatusReq
		updateResponse.CreatedAt = res.CreatedAt
		updateResponse.UpdatedAt = res.UpdatedAt
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
	}
}

// Delete Request Product
func (rpc *ReqProductsController) DeleteReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		success, errQuery := rpc.model.Delete(paramId)
		if !success {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Request Product not found", errQuery))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success delete Request Product", nil))
	}
}

// Searching
func (rpc *ReqProductsController) SearchReqProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		reqProducts, err := rpc.model.SearchReqProduct(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", err))
		}

		var searchResponse []response.ReqProductResponse

		for _, reqProduct := range reqProducts {
			searchResponse = append(searchResponse, response.ReqProductResponse{
				Id:          reqProduct.Id,
				IdEmployee:  reqProduct.IdEmployee,
				ProductName: reqProduct.ProductName,
				Quantity:    reqProduct.Quantity,
				Note:        reqProduct.Note,
				StatusReq:   reqProduct.StatusReq,
				CreatedAt:   reqProduct.CreatedAt,
				UpdatedAt:   reqProduct.UpdatedAt,
				DeletedAt:   reqProduct.DeletedAt,
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", searchResponse))
	}
}
