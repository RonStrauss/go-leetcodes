package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(reverseVowels("Bob: \"Did Anna peep?\" Anna: \"Did Bob?\""))
}

func reverseVowels(s string) string {

	vowels := "AEIOUaeiou"
	ptr1, ptr2 := 0, len(s)-1
	sl := []byte(s)

	for ptr1 <= ptr2 && ptr1 < len(s) && ptr2 > -1 {
		c1, c2 := sl[ptr1], sl[ptr2]
		v1, v2 := strings.Contains(vowels, string(c1)), strings.Contains(vowels, string(c2))
		if !v1 {
			ptr1++
		}
		if !v2 && ptr2 != ptr1 {
			ptr2--
		}

		if v1 && v2 {
			temp := sl[ptr1]
			sl[ptr1] = sl[ptr2]
			sl[ptr2] = temp
			ptr1++
			ptr2--
		}
	}

	return string(sl)
}
