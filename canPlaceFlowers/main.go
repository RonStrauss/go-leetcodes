package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	arr := make([]int, 10_000_000) // Create a slice of length 1,000,000

	for i := range arr {
		if rand.Float64() > 0.5 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}

	log.Println(canPlaceFlowers(arr, rand.Intn(10_000_000)+1))
	log.Println("total time taken:", time.Since(start))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	i := 0
	for i < len(flowerbed) && n > 0 {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i += 2
		} else {
			i++
		}
	}
	var ans bool
	if n == 0 {
		ans = true
	}
	return ans

}
