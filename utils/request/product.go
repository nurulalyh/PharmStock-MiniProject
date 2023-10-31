package request

type ProductRequest struct {
	Name          string `json:"name" form:"name"`
	Image         string `json:"image" form:"image"`
	IdCatProduct  string `json:"id_cat_product" form:"id_cat_product"`
	MfDate        string `json:"mf_date" form:"mf_date"`
	ExpDate       string `json:"exp_date" form:"exp_date"`
	BatchNumber   int    `json:"batch_number" form:"batch_number"`
	UnitPrice     int    `json:"unit_price" form:"unit_price"`
	Stock         int    `json:"stock" form:"stock"`
	IdDistributor string `json:"id_distributor" form:"id_distributor"`
}

type AIDescriptionRequest struct {
	Name string `json:"name" form:"name"`
}
