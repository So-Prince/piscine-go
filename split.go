package piscine

func Split(s, sep string) []string {
	result := []string{}
	start := 0
	seplen := len(sep)

	for i := 0; i < len(s); i++ {
		if seplen > 0 && i+seplen <= len(s) && s[i:i+seplen] == sep {
			// Found full separator
			word := s[start:i]
			result = append(result, word)
			i += seplen - 1 // Skip the entire sep
			start = i + 1
		}
	}

	// Add the last part
	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}
