package main

import (
	"log"
	"slices"
	"strings"
)

func main() {
	log.Println(reverseWords(" a good   example \t\n "))
}

func reverseWords(s string) string {
	words := strings.Fields(s)

	slices.Reverse(words)

	return strings.Join(words, " ")
}
