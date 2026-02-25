package piscine

import "github.com/01-edu/z01"

// PrintNbrInOrder prints digits of an integer in ascending order
func PrintNbrInOrder(n int) {
	// Handle zero directly
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract digits into a slice
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Sort digits in ascending order (simple bubble sort)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Print digits as characters
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
