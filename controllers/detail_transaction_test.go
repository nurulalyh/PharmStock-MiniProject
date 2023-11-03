package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestDetailTransactionsController_CreateDetailTransaction(t *testing.T) {
	tests := []struct {
		name string
		dtc  *DetailTransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dtc.CreateDetailTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailTransactionsController.CreateDetailTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetailTransactionsController_GetAllDetailTransaction(t *testing.T) {
	tests := []struct {
		name string
		dtc  *DetailTransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dtc.GetAllDetailTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailTransactionsController.GetAllDetailTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetailTransactionsController_SearchDetailTransaction(t *testing.T) {
	tests := []struct {
		name string
		dtc  *DetailTransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dtc.SearchDetailTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailTransactionsController.SearchDetailTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
