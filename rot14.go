package piscine

func Rot14(s string) string {
	result := []rune{} // a slice to store the new characters

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			// lowercase letter
			newR := r + 14
			if newR > 'z' {
				newR = newR - 26 // wrap around alphabet
			}
			result = append(result, newR)
		} else if r >= 'A' && r <= 'Z' {
			// uppercase letter
			newR := r + 14
			if newR > 'Z' {
				newR = newR - 26 // wrap around alphabet
			}
			result = append(result, newR)
		} else {
			// not a letter (keep as is)
			result = append(result, r)
		}
	}

	return string(result)
}
