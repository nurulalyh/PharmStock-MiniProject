package response

import (
	"time"
)

type LoginResponse struct {
	Id           string `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
	Token        any    `json:"token" form:"token"`
	RefreshToken any    `json:"refresh_token" form:"refresh_token"`
}

type UsersResponse struct {
	Id           string                `json:"id" form:"id"`
	Name         string                `json:"name" form:"name"`
	Username     string                `json:"username" form:"username"`
	Email        string                `json:"email" form:"email"`
	Phone        string                `json:"phone" form:"phone"`
	Address      string                `json:"address" form:"address"`
	Role         string                `json:"role" form:"role"`
	CreatedAt    time.Time             `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at" form:"updated_at"`
	Transactions []TransactionResponse `json:"transaction_history" form:"transaction_history"`
	ReqProducts  []ReqProductResponse  `json:"request_product_history" form:"request_product_history"`
}
