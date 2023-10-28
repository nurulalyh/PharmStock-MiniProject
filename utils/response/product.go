package response

import (
	"time"

	"gorm.io/gorm"
)

type ProductResponse struct {
	Id            string         `json:"id" form:"id"`
	Name          string         `json:"name" form:"name"`
	Photo         string         `json:"photo" form:"photo"`
	IdCatProduct  string         `json:"id_cat_product" form:"id_cat_product"`
	MfDate        time.Time      `json:"mf_date" form:"mf_date"`
	ExpDate       time.Time      `json:"exp_date" form:"exp_date"`
	BatchNumber   int            `json:"batch_number" form:"batch_number"`
	UnitPrice     int            `json:"unit_price" form:"unit_price"`
	Stock         int            `json:"stock" form:"stock"`
	Description   string         `json:"description" form:"description"`
	IdDistributor string         `json:"id_distributor" form:"id_distributor"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
