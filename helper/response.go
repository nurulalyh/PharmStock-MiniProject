package helper

func ErrorResponse (status string, message any) map[string]any {
	var response = map[string]any{}

	response["status"] = status
	if message != nil {
		response["message"] = message
	}

	return response
}