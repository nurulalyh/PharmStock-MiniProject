package controllers

import (
	"pharm-stock/models"
	"pharm-stock/utils/request"
	"pharm-stock/utils/response"
	"testing"
)

type UsersMockModel struct {
	users []models.Users
}

type LoginMockRequest struct {
	login []request.LoginRequest
}

type UsersMockRequest struct {
	userReq []request.UsersRequest
}

type LoginMockResponse struct {
	login []response.LoginResponse
}

type InsertUsersMockResponse struct {
	userResp []response.InsertUsersResponse
}

func (um *UsersMockModel) Login(username string) (*models.Users, error) {
	data := models.Users{}
	return &data, nil
}

func (um *UsersMockModel) Insert(newUser models.Users) (*models.Users, error) {
	data := models.Users{}
	return &data, nil
}

func (um *UsersMockModel) SelectAll(limit, offset int) ([]models.Users, error) {
	data := []models.Users{}
	return data, nil
}

func (um *UsersMockModel) Update(updatedData models.Users) (*models.Users, error) {
	data := models.Users{}
	return &data, nil
}

func (um *UsersMockModel) Delete(userId string) (bool, error) {
	return true, nil
}

func (um *UsersMockModel) SearchUsers(keyword string, limit int, offset int) ([]models.Users, error) {
	data := []models.Users{}
	return data, nil
}

func (um *UsersMockModel) InsertAdmin(newUser models.Users) (*models.Users, error) {
	data := models.Users{}
	return &data, nil
}

func TestUsersController_Login(t *testing.T) {
	t.Run("login success", func(t *testing.T) {

	})
	t.Run("Invalid user input", func(t *testing.T) {

	})
	t.Run("Wrong password", func(t *testing.T) {

	})
	t.Run("Cannot login", func(t *testing.T) {

	})
	// t.Run("Cannot generate token", func(t *testing.T) {

	// })
}

func TestUsersController_CreateUser(t *testing.T) {
	// t.Run("Missing Role", func(t *testing.T) {

	// })
	t.Run("Don't Have Permission", func(t *testing.T) {

	})
	t.Run("Invalid user input", func(t *testing.T) {

	})
	// t.Run("Cannot hash password", func(t *testing.T) {

	// })
	t.Run("Cannot process data", func(t *testing.T) {

	})
	t.Run("Successfully created user", func(t *testing.T) {

	})
	// t.Run("Unauthorized accessn", func(t *testing.T) {

	// })
}

func TestUsersController_GetAllUsers(t *testing.T) {
	// t.Run("Missing Role", func(t *testing.T) {

	// })
	t.Run("Don't Have Permission", func(t *testing.T) {

	})
	t.Run("Error get all", func(t *testing.T) {

	})
	t.Run("Success get all users", func(t *testing.T) {

	})
	// t.Run("Unauthorized access", func(t *testing.T) {

	// })
}

func TestUsersController_UpdateUser(t *testing.T) {
	// t.Run("Missing Role", func(t *testing.T) {

	// })
	t.Run("Don't Have Permission", func(t *testing.T) {

	})
	t.Run("Invalid user input", func(t *testing.T) {

	})
	t.Run("Cannot process data", func(t *testing.T) {

	})
	t.Run("Successfully updated user", func(t *testing.T) {

	})
	// t.Run("Unauthorized access", func(t *testing.T) {

	// })
}

func TestUsersController_DeleteUser(t *testing.T) {
	// t.Run("Missing Role", func(t *testing.T) {

	// })
	t.Run("Don't Have Permission", func(t *testing.T) {

	})
	t.Run("User not found", func(t *testing.T) {

	})
	t.Run("Successfully deleted user", func(t *testing.T) {

	})
	// t.Run("Unauthorized access", func(t *testing.T) {

	// })
}

func TestUsersController_SearchUsers(t *testing.T) {
	// t.Run("Missing Role", func(t *testing.T) {

	// })
	t.Run("Don't Have Permission", func(t *testing.T) {

	})
	t.Run("Error search", func(t *testing.T) {

	})
	t.Run("Success search", func(t *testing.T) {

	})
	// t.Run("Unauthorized access", func(t *testing.T) {

	// })
}

func TestUsersController_CreateAdmin(t *testing.T) {
	t.Run("Invalid user input", func(t *testing.T) {

	})
	t.Run("Cannot hashed password", func(t *testing.T) {

	})
	t.Run("Cannot process data", func(t *testing.T) {

	})
	t.Run("Successfully created admin", func(t *testing.T) {

	})
}
