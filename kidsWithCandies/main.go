package main

func main() {
	KidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	highestNumber := 0

	for candy := range candies {
		if highestNumber < candy {
			highestNumber = candy
		}
	}
	mappedArray := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= highestNumber {
			mappedArray[i] = true
		}
	}

	return mappedArray
}
