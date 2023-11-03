package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestTransactionsController_CreateTransaction(t *testing.T) {
	tests := []struct {
		name string
		tc   *TransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.CreateTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsController.CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionsController_GetAllTransaction(t *testing.T) {
	tests := []struct {
		name string
		tc   *TransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.GetAllTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsController.GetAllTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionsController_SearchTransaction(t *testing.T) {
	tests := []struct {
		name string
		tc   *TransactionsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.SearchTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsController.SearchTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
