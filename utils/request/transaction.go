package request

type InsertTransactionsRequest struct {
	IdEmployee        string               `json:"id_employee" form:"id_employee"`
	Type              string               `json:"type" form:"type"`
}