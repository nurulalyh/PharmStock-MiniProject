package response

import (
	"time"

	"gorm.io/gorm"
)

type DetailTransactionsResponse struct {
	Id            string         `json:"id" form:"id"`
	IdTransaction string         `json:"id_transaction" form:"id_transaction"`
	IdProduct     string         `json:"id_product" form:"id_product"`
	Quantity      int         `json:"quantity" form:"quantity"`
	Price         int         `json:"price" form:"price"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
