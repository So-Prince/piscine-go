package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// --- Step 1: Validate the base ---
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// --- Step 2: Handle negative numbers ---
	if nbr < 0 {
		z01.PrintRune('-')
		// Convert to positive (watch out for overflow)
		if nbr == -9223372036854775808 {
			PrintNbrBase(922337203685477580, base)
			z01.PrintRune(rune(base[8]))
			return
		}
		nbr = -nbr
	}

	// --- Step 3: Convert number into the base recursively ---
	baseLen := len(base)
	if nbr >= baseLen {
		PrintNbrBase(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}

// Helper function to check if base is valid
func isValidBase(base string) bool {
	// A base must have at least two characters
	if len(base) < 2 {
		return false
	}

	// Create a map to detect duplicates
	seen := make(map[rune]bool)

	for _, ch := range base {
		// A base cannot contain '+' or '-'
		if ch == '+' || ch == '-' {
			return false
		}

		// Each character must be unique
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}
