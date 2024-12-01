package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startrepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		reader.Scan()
		
		words := clearInputs(reader.Text())
		if len(words) == 0 {
			continue
		}

		fmt.Println((words))
	}
}

func clearInputs(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}