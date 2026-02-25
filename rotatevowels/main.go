package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1:]
	str := ""
	for i, arg := range args {
		if i > 0 {
			str += " "
		}
		str += arg
	}

	runes := []rune(str)
	// collect vowel indices
	var vowels []int
	for i, r := range runes {
		if isVowel(r) {
			vowels = append(vowels, i)
		}
	}

	// swap vowels mirrored
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		runes[vowels[i]], runes[vowels[j]] = runes[vowels[j]], runes[vowels[i]]
	}

	// print result
	for _, r := range runes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
