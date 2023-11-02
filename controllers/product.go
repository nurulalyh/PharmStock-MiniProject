package controllers

import (
	"fmt"
	"net/http"
	"os"
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
type ProductsControllerInterface interface {
	CreateProduct() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	SearchProduct() echo.HandlerFunc
}

// Connect into db and model
type ProductsController struct {
	config configs.Config
	model  models.ProductsModelInterface
}

// Create new instance from ProductsController
func NewProductsControllerInterface(m models.ProductsModelInterface) ProductsControllerInterface {
	return &ProductsController{
		model: m,
	}
}

// Create Product
func (pc *ProductsController) CreateProduct() echo.HandlerFunc {
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

			var input = request.ProductRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Product input", errBind.Error()))
			}
	
			file, errImg := c.FormFile("image")
			if errImg != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Image not found, upload image before add product", errImg.Error()))
			}
	
			src, errOpen := file.Open()
			if errOpen != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error open image", errOpen.Error()))
			}
			defer src.Close()
	
			url, _, errUpload := helper.UploadImage(src)
			if errUpload != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error upload image to cloudinary", errUpload.Error()))
			}
	
			userInput := fmt.Sprintf("Deskripsi singkat dalam paragraf mengenai %s dan indikasinya", input.Name)
			apiKey, found := os.LookupEnv("OPENAI_API_KEY")
			if !found {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Open AI API key not found", nil))
			}
	
			generateDescription, errAI := pc.model.AIGenerateDescription(userInput, apiKey)
			if errAI != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error when generate description form AI", errAI.Error()))
			}
	
			var newProduct = models.Products{}
			newProduct.Name = input.Name
			newProduct.Image = url
			newProduct.IdCatProduct = input.IdCatProduct
			newProduct.MfDate = input.MfDate
			newProduct.ExpDate = input.ExpDate
			newProduct.BatchNumber = input.BatchNumber
			newProduct.UnitPrice = input.UnitPrice
			newProduct.Stock = input.Stock
			newProduct.Description = generateDescription
			newProduct.IdDistributor = input.IdDistributor
	
			var res, errQuery = pc.model.Insert(newProduct)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
			}
	
			var insertResponse = response.InsertProductResponse{}
			insertResponse.Id = res.Id
			insertResponse.Name = res.Name
			insertResponse.Image = res.Image
			insertResponse.IdCatProduct = res.IdCatProduct
			insertResponse.MfDate = res.MfDate
			insertResponse.ExpDate = res.ExpDate
			insertResponse.BatchNumber = res.BatchNumber
			insertResponse.UnitPrice = res.UnitPrice
			insertResponse.Stock = res.Stock
			insertResponse.Description = res.Description
			insertResponse.IdDistributor = res.IdDistributor
			insertResponse.CreatedAt = res.CreatedAt
	
			return c.JSON(http.StatusCreated, helper.FormatResponse("success create Product", insertResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Get All Product
func (pc *ProductsController) GetAllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = pc.model.SelectAll(limit, offset)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all Product, ", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all Product, ", res))
	}
}

// Update Product
func (pc *ProductsController) UpdateProduct() echo.HandlerFunc {
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
			var input = models.Products{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid Product input", errBind.Error()))
			}
	
			input.Id = paramId
	
			var res, errQuery = pc.model.Update(input)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
			}
	
			updateResponse := response.UpdateProductResponse{}
			updateResponse.Id = res.Id
			updateResponse.Name = res.Name
			updateResponse.Image = res.Image
			updateResponse.IdCatProduct = res.IdCatProduct
			updateResponse.MfDate = res.MfDate
			updateResponse.ExpDate = res.ExpDate
			updateResponse.BatchNumber = res.BatchNumber
			updateResponse.UnitPrice = res.UnitPrice
			updateResponse.Stock = res.Stock
			updateResponse.Description = res.Description
			updateResponse.IdDistributor = res.IdDistributor
			updateResponse.CreatedAt = res.CreatedAt
			updateResponse.UpdatedAt = res.UpdatedAt
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Searching
func (pc *ProductsController) SearchProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		products, err := pc.model.SearchProduct(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search category products, something happened", err.Error()))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Search category product success", products))
	}
}
