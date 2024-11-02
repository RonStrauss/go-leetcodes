package main

func main() {
	MergeAlternately("goodbye", "world")
}

func MergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	maxLength := len1

	if len2 > maxLength {
		maxLength = len2
	}

	var solution []byte

	for i := 0; i < maxLength; i++ {
		if len1 > i {
			solution = append(solution, word1[i])
		}
		if len2 > i {
			solution = append(solution, word2[i])
		}
	}

	return string(solution)
}
