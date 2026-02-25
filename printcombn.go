package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// Make a list to hold our blocks
	comb := make([]int, n)

	// Start building from position 0
	printCombinations(comb, n, 0)

	// Print newline at end
	z01.PrintRune('\n')
}

func printCombinations(comb []int, n int, pos int) {
	// Are we done building? (Tower is n blocks tall)
	if pos == n {
		printComb(comb, n) // Show the tower!
		return
	}

	// What's the smallest block we can use here?
	start := 0
	if pos > 0 {
		start = comb[pos-1] + 1 // Must be bigger than previous block
	}

	// Try every possible block from start to 9
	// But make sure we have enough blocks left!
	for digit := start; digit <= 9-(n-pos-1); digit++ {
		comb[pos] = digit                 // Place this block
		printCombinations(comb, n, pos+1) // Build next level
	}
}

func printComb(comb []int, n int) {
	// Print all the blocks in the tower
	for i := 0; i < n; i++ {
		z01.PrintRune(rune('0' + comb[i]))
	}

	// Is this the last possible tower?
	isLast := true
	for i := 0; i < n; i++ {
		if comb[i] != 10-n+i {
			isLast = false
			break
		}
	}

	// Add comma and space if not last
	if !isLast {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
