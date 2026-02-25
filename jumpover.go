package piscine

func JumpOver(str string) string {
	result := ""
	count := 0

	for _, ch := range str {
		count++
		if count%3 == 0 {
			result += string(ch)
		}
	}

	return result + "\n"
}
