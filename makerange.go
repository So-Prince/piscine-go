package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min // e.g., 10 - 5 = 5 numbers: 5,6,7,8,9

	// Step 2: Make a slice of exactly that size
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}
