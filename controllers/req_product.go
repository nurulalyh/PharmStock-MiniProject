package controllers

// import (
// 	"net/http"
// 	"pharm-stock/configs"
// 	"pharm-stock/helper"
// 	"pharm-stock/models"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type ReqProductControllerInterface interface {
// 	CreateReqProduct() echo.HandlerFunc
// 	GetAllReqProduct() echo.HandlerFunc
// 	GetReqProductById() echo.HandlerFunc
// 	UpdateReqProduct() echo.HandlerFunc
// 	DeleteReqProduct() echo.HandlerFunc
// }

// type ReqProductController struct {
// 	config configs.Config
// 	model  models.ReqProductModelInterface
// }

// func NewReqProductControllerInterface(m models.ReqProductModelInterface) ReqProductControllerInterface {
// 	return &ReqProductController{
// 		model: m,
// 	}
// }

// func (rpc *ReqProductController) CreateReqProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input = models.ReqProduct{}
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", nil))
// 		}

// 		var res = rpc.model.Insert(input)
// 		if res == nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
// 		}

// 		return c.JSON(http.StatusCreated, helper.FormatResponse("success create Request Product", res))
// 	}
// }

// func (rpc *ReqProductController) GetAllReqProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var res = rpc.model.SelectAll()

// 		if res == nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Request Product, ", nil))
// 		}

// 		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Request Product, ", res))
// 	}
// }

// func (rpc *ReqProductController) GetReqProductById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var paramId = c.Param("id")

// 		cnv, err := strconv.Atoi(paramId)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
// 		}

// 		var res = rpc.model.SelectById(cnv)
// 		if res == nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get Request Product by id, ", nil))
// 		}

// 		return c.JSON(http.StatusOK, helper.FormatResponse("Success get Request Product", res))
// 	}
// }

// func (rpc *ReqProductController) UpdateReqProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var paramId = c.Param("id")
// 		cnv, err := strconv.Atoi(paramId)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
// 		}

// 		var input = models.ReqProduct{}
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Request Product input", nil))
// 		}

// 		input.Id = cnv

// 		var res = rpc.model.Update(input)
// 		if res == nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
// 		}

// 		return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", res))
// 	}
// }

// func (rpc *ReqProductController) DeleteReqProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 	  var paramId = c.Param("id")
  
// 	  cnv, err := strconv.Atoi(paramId)
// 	  if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
// 	  }
  
// 	  success := rpc.model.Delete(cnv)
// 	  if !success {
// 		return c.JSON(http.StatusNotFound, helper.FormatResponse("Request Product not found", nil))
// 	  }
  
// 	  return c.JSON(http.StatusOK, helper.FormatResponse("Success delete Request Product", nil))
// 	}
// }