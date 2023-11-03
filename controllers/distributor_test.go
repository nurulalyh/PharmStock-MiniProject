package controllers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestDistributorsController_CreateDistributor(t *testing.T) {
	tests := []struct {
		name string
		dc   *DistributorsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.CreateDistributor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributorsController.CreateDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistributorsController_GetAllDistributor(t *testing.T) {
	tests := []struct {
		name string
		dc   *DistributorsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.GetAllDistributor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributorsController.GetAllDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistributorsController_UpdateDistributor(t *testing.T) {
	tests := []struct {
		name string
		dc   *DistributorsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.UpdateDistributor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributorsController.UpdateDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistributorsController_DeleteDistributor(t *testing.T) {
	tests := []struct {
		name string
		dc   *DistributorsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.DeleteDistributor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributorsController.DeleteDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistributorsController_SearchDistributor(t *testing.T) {
	tests := []struct {
		name string
		dc   *DistributorsController
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.SearchDistributor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributorsController.SearchDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}
