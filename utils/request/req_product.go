package request

type InsertReqProductRequest struct {
	IdEmployee  string `json:"id_employee" form:"id_employee"`
	ProductName string `json:"product_name" form:"product_name"`
	Quantity    int    `json:"quantity" form:"quantity"`
	Note        string `json:"note" form:"note"`
}

type UpdateReqProductRequest struct {
	Id        string `json:"id" form:"id"`
	StatusReq string `json:"status_req" form:"status_req"`
}
