package main

import (
	"os"

	"github.com/01-edu/z01"
)

func putRune(r rune) {
	z01.PrintRune(r)
}

func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	n := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}

	return n * sign, true
}

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	printed := false

	for _, a := range args {
		n, ok := atoi(a)

		if !ok || n < 1 || n > 26 {
			putRune(' ')
			printed = true
			continue
		}

		r := rune('a' + n - 1)

		if upper {
			r -= 32
		}

		putRune(r)
		printed = true
	}

	if printed {
		putRune('\n')
	}
}
