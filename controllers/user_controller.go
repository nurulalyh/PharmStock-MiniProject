package controllers

import (
	"net/http"
	"pharm-stock/configs"
	"pharm-stock/helper"
	"pharm-stock/models"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	config configs.Config
	model  models.UsersModel
}

func (uc *UserController) InitUserController(um models.UsersModel, cfg configs.Config) {
	uc.config = cfg
	uc.model = um
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

		return c.JSON(http.StatusOK, helper.FormatResponse("success", jwtToken))
	}
}
