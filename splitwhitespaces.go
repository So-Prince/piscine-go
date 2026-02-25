package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	current := ""

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	// Don't forget the last word!
	if current != "" {
		words = append(words, current)
	}

	return words
}
