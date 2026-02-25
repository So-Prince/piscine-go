package main

import (
	"fmt"
	"os"
)

func main() {
	for _, str := range os.Args[1:] {
		if str == "01" || str == "galaxy" || str == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
