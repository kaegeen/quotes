package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"The only way to do great work is to love what you do. – Steve Jobs",
		"Success is not the key to happiness. Happiness is the key to success. – Albert Schweitzer",
		"Don’t watch the clock; do what it does. Keep going. – Sam Levenson",
		"You miss 100% of the shots you don’t take. – Wayne Gretzky",
		"The best way to predict the future is to create it. – Peter Drucker",
		"It always seems impossible until it’s done. – Nelson Mandela",
		"Hardships often prepare ordinary people for an extraordinary destiny. – C.S. Lewis",
		"Believe you can and you're halfway there. – Theodore Roosevelt",
		"Act as if what you do makes a difference. It does. – William James",
		"Keep your face always toward the sunshine—and shadows will fall behind you. – Walt Whitman",
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random Inspirational Quote Generator")
	fmt.Println("=====================================")
	fmt.Println("Type 'next' to get a new quote or 'exit' to quit.\n")

	for {
		fmt.Print("Command (next/exit): ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "next":
			randomIndex := rand.Intn(len(quotes))
			fmt.Printf("\nQuote: %s\n\n", quotes[randomIndex])
		case "exit":
			fmt.Println("Thanks for using the quote generator. Stay inspired!")
			return
		default:
			fmt.Println("Invalid command. Please type 'next' or 'exit'.")
		}
	}
}
