package response

import (
	"time"

	"gorm.io/gorm"
)

type ReqProductResponse struct {
	Id          string         `json:"id" form:"id"`
	IdEmployee  string         `json:"id_employee" form:"id_employee"`
	ProductName string         `json:"product_name" form:"product_name"`
	Quantity    int            `json:"quantity" form:"quantity"`
	Note        string         `json:"note" form:"note"`
	StatusReq   string         `json:"status" form:"status"`
	CreatedAt   time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
