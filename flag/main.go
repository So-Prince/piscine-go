package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	printStr(`--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`)
}

func orderString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""
	help := false
	skipNext := false

	for i := 0; i < len(args); i++ {
		if skipNext {
			skipNext = false
			continue
		}

		arg := args[i]

		if arg == "--help" || arg == "-h" {
			help = true
			break
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else if arg == "-i" && i+1 < len(args) {
			insert = args[i+1]
			skipNext = true
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if len(arg) > 0 && arg[0] != '-' && mainStr == "" {
			mainStr = arg
		}
	}

	if help {
		printHelp()
		return
	}

	result := mainStr + insert

	if order {
		result = orderString(result)
	}

	printStr(result)
}
