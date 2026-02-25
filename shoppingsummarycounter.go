package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	currentItem := ""

	if len(str) == 0 {
		counts[""] = 1
		return counts
	}

	for _, r := range str {
		if r == ' ' {
			counts[currentItem]++
			currentItem = ""
		} else {
			currentItem += string(r)
		}
	}

	counts[currentItem]++

	return counts
}
