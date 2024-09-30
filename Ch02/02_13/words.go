package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	fmt.Println(text)

	text = strings.ToLower(text)

	count := map[string]int{}

	words := strings.Fields(text)

	for _, word := range words {
		count[word]++
	}

	fmt.Println(count)
}
