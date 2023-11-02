package response

import (
	"time"
)

type InsertDetailTransactionsResponse struct {
	Id            string    `json:"id" form:"id"`
	IdTransaction string    `json:"id_transaction" form:"id_transaction"`
	IdProduct     string    `json:"id_product" form:"id_product"`
	Quantity      int       `json:"quantity" form:"quantity"`
	Price         int       `json:"price" form:"price"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}
