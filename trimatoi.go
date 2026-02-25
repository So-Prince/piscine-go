package piscine

func TrimAtoi(s string) int {
	sign := 1   // Default: positive
	result := 0 // The integer we’ll build
	foundNumber := false

	for _, char := range s {
		if char == '-' && !foundNumber {
			sign = -1 // Negative sign before digits
		}
		if char >= '0' && char <= '9' {
			foundNumber = true
			result = result*10 + int(char-'0') // Build number
		}
	}

	return result * sign
}
