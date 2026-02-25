package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 6 * 7
	ptr.y = 3 * 7
}

func printStr() {
	s := []rune{
		'x', ' ', '=', ' ',
		'0' + 4, '0' + 2,
		',', ' ',
		'y', ' ', '=', ' ',
		'0' + 2, '0' + 1,
		'\n',
	}
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	printStr()
}
