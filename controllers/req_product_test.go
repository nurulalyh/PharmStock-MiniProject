package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestReqProductsController_CreateReqProduct(t *testing.T) {
	tests := []struct {
		name string
		rpc  *ReqProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rpc.CreateReqProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqProductsController.CreateReqProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqProductsController_GetAllReqProduct(t *testing.T) {
	tests := []struct {
		name string
		rpc  *ReqProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rpc.GetAllReqProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqProductsController.GetAllReqProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqProductsController_UpdateReqProduct(t *testing.T) {
	tests := []struct {
		name string
		rpc  *ReqProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rpc.UpdateReqProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqProductsController.UpdateReqProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqProductsController_SearchReqProduct(t *testing.T) {
	tests := []struct {
		name string
		rpc  *ReqProductsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rpc.SearchReqProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqProductsController.SearchReqProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
