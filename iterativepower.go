package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	// Any number to power 0 is 1
	if power == 0 {
		return 1
	}

	result := 1

	// Multiply nb by itself power times
	for i := 0; i < power; i++ {
		result = result * nb
	}

	return result
}
