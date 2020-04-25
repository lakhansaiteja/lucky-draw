package lucky_draw

import (
	"fmt"
	"time"
)

func RemoveEmptyItems(items []string) []string {
	var itemsModified []string
	for _, v := range items {
		if v != "" {
			itemsModified = append(itemsModified, v)
		}
	}
	return itemsModified
}

func PrintTheWinner(luckyNumber int, itemsModified []string) {
	fmt.Printf("\n")
	fmt.Print("The lucky winner is...\n\n")
	time.Sleep(1 * time.Second)
	fmt.Print("*drum roll*\n\n")
	time.Sleep(2 * time.Second)
	fmt.Println(fmt.Sprintf("%s%s", itemsModified[luckyNumber], "!"))
}
