package helpers

func ErrorResponse(message string) map[string]any {
	return map[string]any{
		"status":  "failed",
		"message": message,
	}
}

func SuccesResponses(message string) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
	}
}

func SuccessWithDataResponse(message string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}