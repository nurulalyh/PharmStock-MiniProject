package request

import "time"

type InsertProductRequest struct {
	Name          string    `json:"name" form:"name"`
	Photo         string    `json:"photo" form:"photo"`
	IdCatProduct  string    `json:"id_cat_product" form:"id_cat_product"`
	MfDate        time.Time `json:"mf_date" form:"mf_date"`
	ExpDate       time.Time `json:"exp_date" form:"exp_date"`
	BatchNumber   int       `json:"batch_number" form:"batch_number"`
	UnitPrice     int       `json:"unit_price" form:"unit_price"`
	Stock         int       `json:"stock" form:"stock"`
	Description   string    `json:"description" form:"description"`
	IdDistributor string    `json:"id_distributor" form:"id_distributor"`
}
