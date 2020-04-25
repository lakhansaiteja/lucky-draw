package main

import (
	"log"
	"lucky_draw"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file was provided as an argument to the script. Please provide the intended file as an argument.")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error in the path of file provided.")
	}

	var list string
	buf := make([]byte, 1000)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		list = string(buf[:n])
	}

	if list == "" {
		log.Fatal("Error picking winner since the file is empty.")
	}

	items := strings.Split(list, "\n")

	itemsModified := lucky_draw.RemoveEmptyItems(items)
	n := len(itemsModified)
	if n == 0 {
		log.Fatal("Error picking winner since the file is empty.")
	}

	rand.Seed(time.Now().UnixNano())
	luckyNumber := rand.Intn(n)

	lucky_draw.PrintTheWinner(luckyNumber, itemsModified)
}
