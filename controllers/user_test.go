package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUsersController_Login(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.Login(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_CreateUser(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.CreateUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_GetAllUsers(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.GetAllUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_UpdateUser(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.UpdateUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_DeleteUser(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.DeleteUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_SearchUsers(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.SearchUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.SearchUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersController_CreateAdmin(t *testing.T) {
	tests := []struct {
		name string
		uc   *UsersController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uc.CreateAdmin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersController.CreateAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
