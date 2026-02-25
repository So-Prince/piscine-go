package piscine

func intToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var result []byte

	for n > 0 {
		remainder := n % 2
		result = append(result, byte(remainder)+'0')
		n /= 2
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func ActiveBits(n int) int {
	binaryStr := intToBinary(n)

	slice := []byte(binaryStr)
	counter := 0
	for _, bit := range slice {
		if bit == '1' {
			counter++
		}
	}
	return counter
}
