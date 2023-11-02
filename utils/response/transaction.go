package response

import (
	"time"
)

type InsertTransactionResponse struct {
	Id            string    `json:"id" form:"id"`
	IdEmployee    string    `json:"id_employee" form:"id_employee"`
	TotalQuantity int       `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int       `json:"total_price" form:"total_price"`
	Type          string    `json:"type" form:"type"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}
