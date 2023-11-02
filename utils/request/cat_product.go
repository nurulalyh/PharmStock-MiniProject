package request

type CategoryProductRequest struct {
	Name     string `json:"name" form:"name"`
}