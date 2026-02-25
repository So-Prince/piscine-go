package piscine

func BasicJoin(elems []string) string {
	// Step 1: Create an empty string to store the result
	result := ""

	// Step 2: Loop through each string in the slice
	for i := 0; i < len(elems); i++ {
		// Step 3: Add the current string to the result
		result += elems[i]
	}

	// Step 4: Return the final concatenated string
	return result
}
