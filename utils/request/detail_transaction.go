package request

type DetailTransactionsRequest struct {
	IdTransaction string `json:"id_transaction" form:"id_transaction"`
	IdProduct     string `json:"id_product" form:"id_product"`
	Quantity      int `json:"quantity" form:"quantity"`
}