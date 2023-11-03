package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCatProductsController_CreateCatProduct(t *testing.T) {
	tests := []struct {
		name string
		cpc  *CatProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpc.CreateCatProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatProductsController.CreateCatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatProductsController_GetAllCatProduct(t *testing.T) {
	tests := []struct {
		name string
		cpc  *CatProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpc.GetAllCatProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatProductsController.GetAllCatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatProductsController_UpdateCatProduct(t *testing.T) {
	tests := []struct {
		name string
		cpc  *CatProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpc.UpdateCatProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatProductsController.UpdateCatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatProductsController_DeleteCatProduct(t *testing.T) {
	tests := []struct {
		name string
		cpc  *CatProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpc.DeleteCatProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatProductsController.DeleteCatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatProductsController_SearchCatProduct(t *testing.T) {
	tests := []struct {
		name string
		cpc  *CatProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpc.SearchCatProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatProductsController.SearchCatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
