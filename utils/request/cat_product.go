package request

type InsertCategoryProductRequest struct {
	Name     string `json:"name" form:"name"`
}