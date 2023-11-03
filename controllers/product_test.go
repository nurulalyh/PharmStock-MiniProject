package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestProductsController_CreateProduct(t *testing.T) {
	tests := []struct {
		name string
		pc   *ProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.CreateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsController.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsController_GetAllProduct(t *testing.T) {
	tests := []struct {
		name string
		pc   *ProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.GetAllProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsController.GetAllProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsController_UpdateProduct(t *testing.T) {
	tests := []struct {
		name string
		pc   *ProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.UpdateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsController.UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsController_SearchProduct(t *testing.T) {
	tests := []struct {
		name string
		pc   *ProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SearchProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsController.SearchProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
