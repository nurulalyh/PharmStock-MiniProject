package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper/authentication"
	"pharm-stock/models"
	"pharm-stock/utils/request"
	"pharm-stock/utils/response"

	"strconv"

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
			return c.JSON(http.StatusBadRequest, response.FormatResponse("Invalid user input", nil))
		}

		var res, errQuery = uc.model.Login(input.Username, input.Password)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot login, something happend", errQuery.Error()))
		}
		if res.Id == "" {
			return c.JSON(http.StatusNotFound, response.FormatResponse("Data not found", errQuery.Error()))
		}

		errAuth := authentication.ComparePassword(res.Password, input.Password)
		if errAuth != nil {
			return c.JSON(http.StatusUnauthorized, response.FormatResponse("Wrong password", errAuth))
		}

		var jwtToken, errJWT = authentication.GenerateJWT(uc.config.Secret, uc.config.RefreshSecret, res.Id, res.Username, res.Role)
		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot generate token, something happend", errJWT))
		}

		var info = response.LoginResponse{}
		info.Id = res.Id
		info.Name = res.Name
		info.Username = res.Username
		info.Role = res.Role
		if token, exists := jwtToken["access_token"]; exists {
			info.Token = token
		}
		if refreshToken, exists := jwtToken["access_token"]; exists {
			info.RefreshToken = refreshToken
		}

		return c.JSON(http.StatusOK, response.FormatResponse("login success", info))
	}
}

// Create User
func (uc *UsersController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.InsertUsersRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid user input", errBind))
		}

		hashedPassword, err := authentication.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happened", err))
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
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery))
		}

		var insertResponse = response.UsersResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.Username = res.Username
		insertResponse.Password = res.Password
		insertResponse.Email = res.Email
		insertResponse.Phone = res.Phone
		insertResponse.Address = res.Address
		insertResponse.Role = res.Role
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create user", insertResponse))
	}
}

// Get All User
func (uc *UsersController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))

		var res, err = uc.model.SelectAll(limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Error get all users", err.Error()))
		}

		getAllResponse := response.UsersResponse{}
		for _, user := range res{
			getAllResponse.Id = user.Id
			getAllResponse.Name = user.Name
			getAllResponse.Username = user.Username
			getAllResponse.Password = user.Password
			getAllResponse.Email = user.Email
			getAllResponse.Phone = user.Phone
			getAllResponse.Address = user.Address
			getAllResponse.Role = user.Role
			getAllResponse.CreatedAt = user.CreatedAt
			getAllResponse.UpdatedAt = user.UpdatedAt
			getAllResponse.DeletedAt = user.DeletedAt
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success get all users", getAllResponse))
	}
}

// Update User
func (uc *UsersController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		var input = models.Users{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid user input", errBind))
		}

		input.Id = paramId

		var res, errQuery = uc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("cannot process data, something happend", errQuery))
		}

		updateResponse := response.UsersResponse{}
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
		updateResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusOK, response.FormatResponse("Success update data", updateResponse))
	}
}

// Delete User
func (uc *UsersController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		success, errQuery := uc.model.Delete(paramId)
		if !success {
			return c.JSON(http.StatusNotFound, response.FormatResponse("User not found", errQuery))
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Success delete user", nil))
	}
}

// Searching
func (uc *UsersController) SearchUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		users, err := uc.model.SearchUsers(keyword, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot search users, something happened", err))
		}

		searchResponse := response.UsersResponse{}
		for _, user := range users{
			searchResponse.Id = user.Id
			searchResponse.Name = user.Name
			searchResponse.Username = user.Username
			searchResponse.Password = user.Password
			searchResponse.Email = user.Email
			searchResponse.Phone = user.Phone
			searchResponse.Address = user.Address
			searchResponse.Role = user.Role
			searchResponse.CreatedAt = user.CreatedAt
			searchResponse.UpdatedAt = user.UpdatedAt
			searchResponse.DeletedAt = user.DeletedAt
		}

		return c.JSON(http.StatusOK, response.FormatResponse("Search users success", searchResponse))
	}
}

// Create Admin
func (uc *UsersController) CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = request.InsertUsersRequest{}
		if errBind := c.Bind(&input); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.FormatResponse("invalid user input", errBind))
		}

		hashedPassword, err := authentication.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happened", err))
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
			return c.JSON(http.StatusInternalServerError, response.FormatResponse("Cannot process data, something happend", errQuery.Error()))
		}

		var insertResponse = response.UsersResponse{}
		insertResponse.Id = res.Id
		insertResponse.Name = res.Name
		insertResponse.Username = res.Username
		insertResponse.Password = res.Password
		insertResponse.Email = res.Email
		insertResponse.Phone = res.Phone
		insertResponse.Address = res.Address
		insertResponse.Role = res.Role
		insertResponse.CreatedAt = res.CreatedAt
		insertResponse.UpdatedAt = res.UpdatedAt
		insertResponse.DeletedAt = res.DeletedAt

		return c.JSON(http.StatusCreated, response.FormatResponse("success create user", insertResponse))
	}
}
