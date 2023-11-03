package controllers

import (
	"net/http"
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
type UsersControllerInterface interface {
	Login() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	SearchUsers() echo.HandlerFunc
	CreateAdmin() echo.HandlerFunc
}

// Connect into db and model
type UsersController struct {
	config configs.Config
	model  models.UsersModelInterface
}

// Create new instance from UserController
func NewUsersControlInterface(m models.UsersModelInterface) UsersControllerInterface {
	return &UsersController{
		model: m,
	}
}

// Login
func (uc *UsersController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.LoginRequest{}

		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", errBind.Error()))
		}

		var res, errQuery = uc.model.Login(input.Username)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot login, something happend", errQuery.Error()))
		}

		errAuth := authentication.ComparePassword(res.Password, input.Password)
		if errAuth != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Wrong password", errAuth.Error()))
		}

		var jwtToken, errJWT = authentication.GenerateToken(res.Id, res.Username, res.Role)
		if jwtToken == "" {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot generate token, something happend", errJWT.Error()))
		}

		var info = response.LoginResponse{}
		info.Id = res.Id
		info.Name = res.Name
		info.Username = res.Username
		info.Role = res.Role
		info.Token = jwtToken

		return c.JSON(http.StatusOK, helper.FormatResponse("login success", info))
	}
}

// Create User
func (uc *UsersController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, _ := authentication.ExtractToken(userToken)

			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			var input = request.UsersRequest{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", errBind.Error()))
			}

			hashedPassword, errHash := authentication.HashPassword(input.Password)
			if errHash != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot hash password, something happened", errHash.Error()))
			}

			var newUser = models.Users{}
			newUser.Name = input.Name
			newUser.Username = input.Username
			newUser.Password = hashedPassword
			newUser.Email = input.Email
			newUser.Phone = input.Phone
			newUser.Address = input.Address
			
			var res, errQuery = uc.model.Insert(newUser)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happened", errQuery.Error()))
			}

			var insertResponse = response.InsertUsersResponse{}
			insertResponse.Id = res.Id
			insertResponse.Name = res.Name
			insertResponse.Username = res.Username
			insertResponse.Email = res.Email
			insertResponse.Phone = res.Phone
			insertResponse.Address = res.Address
			insertResponse.Role = res.Role
			insertResponse.CreatedAt = res.CreatedAt

			return c.JSON(http.StatusCreated, helper.FormatResponse("Successfully created user", insertResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Get All User
func (uc *UsersController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, _ := authentication.ExtractToken(userToken)
			
			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			limit, _ := strconv.Atoi(c.QueryParam("limit"))
			offset, _ := strconv.Atoi(c.QueryParam("offset"))
	
			var res, errQuery = uc.model.SelectAll(limit, offset)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all users", errQuery.Error()))
			}
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success get all users", res))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Update User
func (uc *UsersController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, _ := authentication.ExtractToken(userToken)

			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}

			var paramId = c.Param("id")
			var input = models.Users{}
			if errBind := c.Bind(&input); errBind != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", errBind.Error()))
			}
	
			input.Id = paramId
	
			var res, errQuery = uc.model.Update(input)
			if res == nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", errQuery.Error()))
			}
	
			updateResponse := response.UpdateUsersResponse{}
			updateResponse.Id = res.Id
			updateResponse.Name = res.Name
			updateResponse.Username = res.Username
			updateResponse.Password = res.Password
			updateResponse.Email = res.Email
			updateResponse.Phone = res.Phone
			updateResponse.Address = res.Address
			updateResponse.Role = res.Role
			updateResponse.CreatedAt = res.CreatedAt
			updateResponse.UpdatedAt = res.UpdatedAt
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", updateResponse))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Delete User
func (uc *UsersController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, _ := authentication.ExtractToken(userToken)
			
			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}
			var paramId = c.Param("id")
	
			success, errQuery := uc.model.Delete(paramId)
			if !success {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("User not found", errQuery.Error()))
			}
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Success delete user", nil))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))
	}
}

// Searching
func (uc *UsersController) SearchUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)

		if userToken != nil && userToken.Valid {
			tokenData, _ := authentication.ExtractToken(userToken)
	
			role, ok := tokenData["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Role information missing in the token", nil))
			}

			if role != "administrator" {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
			}
			
			keyword := c.QueryParam("keyword")
			limit, _ := strconv.Atoi(c.QueryParam("limit"))
			offset, _ := strconv.Atoi(c.QueryParam("offset"))
			users, errQuery := uc.model.SearchUsers(keyword, limit, offset)
			if errQuery != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot search users, something happened", errQuery.Error()))
			}
	
			return c.JSON(http.StatusOK, helper.FormatResponse("Search users success", users))
		}
		return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized access", nil))

	}
}

// Create Admin
func (uc *UsersController) CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.UsersRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", errBind))
		}

		hashedPassword, err := authentication.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot hashed password, something happened", err))
		}

		var newUser = models.Users{}
		newUser.Name = input.Name
		newUser.Username = input.Username
		newUser.Password = hashedPassword
		newUser.Email = input.Email
		newUser.Phone = input.Phone
		newUser.Address = input.Address

		var res, errQuery = uc.model.InsertAdmin(newUser)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.InsertUsersResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.Username = res.Username
		insertResponse.Email = res.Email
		insertResponse.Phone = res.Phone
		insertResponse.Address = res.Address
		insertResponse.Role = res.Role
		insertResponse.CreatedAt = res.CreatedAt

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create administrator", insertResponse))
	}
}
