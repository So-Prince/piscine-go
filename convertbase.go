package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Parse to decimal
	decimal := 0
	for i := 0; i < len(nbr); i++ {
		digit := nbr[i]
		pos := 0
		for j := 0; j < len(baseFrom); j++ {
			if baseFrom[j] == digit {
				pos = j
				break
			}
		}
		decimal = decimal*len(baseFrom) + pos
	}

	// Step 2: Convert decimal to baseTo
	if decimal == 0 {
		return "0"
	}

	result := ""
	for decimal > 0 {
		digit := decimal % len(baseTo)
		result = string(baseTo[digit]) + result
		decimal /= len(baseTo)
	}

	return result
}
