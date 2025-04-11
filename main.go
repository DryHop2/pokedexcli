package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	loweredWords := strings.Fields(strings.ToLower(text))
	return loweredWords
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}
		fmt.Println("Your command was:", words[0])

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
