package helpers

func SuccesResponse(message string, data any) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":   data,
	}
}

func SuccesResponseAuth(user_id, name, token any) map[string]interface{} {
	return map[string]interface{}{
		"user_id": user_id,
		"name":   name,
		"token": token,
	}
}

func ErrorResponse(message string, status string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
	}
}


