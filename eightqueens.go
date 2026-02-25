package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	const size = 8
	var board [size]int // board[col] = row (1-based)

	// check if we can safely place a queen
	var isSafe func(col, row int) bool
	isSafe = func(col, row int) bool {
		for prevCol := 0; prevCol < col; prevCol++ {
			prevRow := board[prevCol]

			// ❌ Same row check
			if prevRow == row {
				return false
			}

			// ❌ Same diagonal check (difference in col == difference in row)
			if abs(prevRow-row) == abs(prevCol-col) {
				return false
			}
		}
		return true
	}

	// recursive backtracking solver
	var solve func(col int)
	solve = func(col int) {
		if col == size {
			// ✅ all queens placed safely → print the board
			for i := 0; i < size; i++ {
				z01.PrintRune(rune(board[i]) + '0')
			}
			z01.PrintRune('\n')
			return
		}

		// try placing queen in each row (1–8)
		for row := 1; row <= size; row++ {
			if isSafe(col, row) {
				board[col] = row
				solve(col + 1) // recurse to next column
			}
		}
	}

	solve(0)
}

// helper absolute value function
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
