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
	t := strings.Split(strings.TrimSpace(s), " ")
	filtered := make([]string, 0)

	for _, w := range t {
		if w != "" && w != " " {
			filtered = append(filtered, strings.TrimSpace(w))
		}
	}

	slices.Reverse(filtered)

	return strings.Join(filtered, " ")

}
