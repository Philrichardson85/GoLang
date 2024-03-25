package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The feild is required",
		"email":    "The feild must have a valid email address",
		"min":      "The feild is too short",
		"max":      "The feild is too long",
	}
}