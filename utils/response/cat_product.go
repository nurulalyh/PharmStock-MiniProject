package response

import (
	"time"
)

type InsertCatProductsResponse struct {
	Id        string    `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type UpdateCatProductsResponse struct {
	Id        string    `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
