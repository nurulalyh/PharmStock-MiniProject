package helper

type Response struct {
	Message string `json:"message" form:"message"`
	Data    any `json:"data" form:"data"`
}

func FormatResponse(message string, data any) *Response {
	var response = Response{}

	response.Message = message
	if data != nil {
		response.Data = data
	}

	return &response
}