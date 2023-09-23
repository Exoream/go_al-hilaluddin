package helper

func SuccesResponse(message string, data any) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"users":   data,
	}
}

func ErrorResponse(message string, status string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
	}
}
