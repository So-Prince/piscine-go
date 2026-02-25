package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	length := 0
	for range deck {
		length++
	}

	cardsPerPlayer := length / 4

	playerNum := 1
	index := 0

	for playerNum <= 4 {
		fmt.Printf("Player %d: ", playerNum)

		count := 0
		for count < cardsPerPlayer {
			fmt.Printf("%d", deck[index])
			index++
			count++

			if count < cardsPerPlayer {
				fmt.Printf(", ")
			}
		}
		z01.PrintRune('\n')
		playerNum++
	}
}
