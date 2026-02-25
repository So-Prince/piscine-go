package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var i rune
	var j rune
	for i = 0; i <= 98; i++ {
		for j = i + 1; j <= 99; j++ {
			z01.PrintRune('0' + i/10)
			z01.PrintRune('0' + i%10)
			z01.PrintRune(' ')
			z01.PrintRune('0' + j/10)
			z01.PrintRune('0' + j%10)
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
