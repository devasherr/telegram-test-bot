package utils

func MaskPhone(phone string) string {
	// TODO: assuming phone number has 10 digits (make sure of that)
	if len(phone) < 10 {
		return phone
	}

	return phone[:7] + "****"
}
