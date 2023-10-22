package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	CreateUser() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
	GetUserById() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UserController struct {
	config configs.Config
	model  models.UserModelInterface
}

func NewUserControlInterface(m models.UserModelInterface) UserControllerInterface {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = models.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var res = uc.model.Insert(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success create user", res))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = models.Login{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var res = uc.model.Login(input.Username, input.Password)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
		}

		if res.Id == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Data not found", nil))
		}

		var jwtToken = helper.GenerateJWT(uc.config.Secret, uc.config.RefreshSecret, res.Id, res.Username, res.Role)

		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data, something happend", nil))
		}

		var info = map[string]any{}
		info["name"] = res.Name
		info["username"] = res.Username
		info["role"] = res.Role

		jwtToken["info"] = info

		return c.JSON(http.StatusOK, helper.FormatResponse("login success", jwtToken))
	}
}

func (uc *UserController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = uc.model.SelectAll()

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get all users, ", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all users, ", res))
	}
}

func (uc *UserController) GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
		}

		var res = uc.model.SelectById(cnv)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Error get user by id, ", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get user", res))
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
		}

		var input = models.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		input.Id = cnv

		var res = uc.model.Update(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", res))
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
	  var paramId = c.Param("id")
  
	  cnv, err := strconv.Atoi(paramId)
	  if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
	  }
  
	  success := uc.model.Delete(cnv)
	  if !success {
		return c.JSON(http.StatusNotFound, helper.FormatResponse("User not found", nil))
	  }
  
	  return c.JSON(http.StatusOK, helper.FormatResponse("Success delete user", nil))
	}
}