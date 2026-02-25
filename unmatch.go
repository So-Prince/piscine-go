package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, num := range a {
		counts[num]++
	}

	for _, num := range a {
		if counts[num]%2 != 0 {
			return num
		}
	}

	return -1
}
