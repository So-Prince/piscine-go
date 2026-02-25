package piscine

func AtoiBase(s string, base string) int {
	// Check if base is valid
	if len(base) < 2 {
		return 0
	}

	// Check if all characters in base are unique and the base doesn't contain '+' or '-'
	baseMap := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return 0
		}
		if baseMap[char] {
			return 0
		}
		baseMap[char] = true
	}

	// Convert the string to an integer based on the base
	result := 0
	baseLength := len(base)
	for _, char := range s {
		// Find the index of the character in the base string
		index := -1
		for i, baseChar := range base {
			if baseChar == char {
				index = i
				break
			}
		}

		// If character is not found in base, return 0
		if index == -1 {
			return 0
		}

		// Update the result by multiplying the current value by the base and adding the current digit's value
		result = result*baseLength + index
	}

	return result
}
