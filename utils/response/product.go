package response

import (
	"time"
)

type InsertProductResponse struct {
	Id            string    `json:"id" form:"id"`
	Name          string    `json:"name" form:"name"`
	Image         string    `json:"image" form:"image"`
	IdCatProduct  string    `json:"id_cat_product" form:"id_cat_product"`
	MfDate        string    `json:"mf_date" form:"mf_date"`
	ExpDate       string    `json:"exp_date" form:"exp_date"`
	BatchNumber   int       `json:"batch_number" form:"batch_number"`
	UnitPrice     int       `json:"unit_price" form:"unit_price"`
	Stock         int       `json:"stock" form:"stock"`
	Description   string    `json:"description" form:"description"`
	IdDistributor string    `json:"id_distributor" form:"id_distributor"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}

type UpdateProductResponse struct {
	Id            string    `json:"id" form:"id"`
	Name          string    `json:"name" form:"name"`
	Image         string    `json:"image" form:"image"`
	IdCatProduct  string    `json:"id_cat_product" form:"id_cat_product"`
	MfDate        string    `json:"mf_date" form:"mf_date"`
	ExpDate       string    `json:"exp_date" form:"exp_date"`
	BatchNumber   int       `json:"batch_number" form:"batch_number"`
	UnitPrice     int       `json:"unit_price" form:"unit_price"`
	Stock         int       `json:"stock" form:"stock"`
	Description   string    `json:"description" form:"description"`
	IdDistributor string    `json:"id_distributor" form:"id_distributor"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

type AIDescriptionResponse struct {
	Description string `json:"description" form:"description"`
}
