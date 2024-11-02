package main

import (
	"log"
)

func main() {
	log.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	solution := make([]int, len(nums))
	if len(nums) > 1 {
		solution[0] = productOfAll(nums[1:])
	}

	for i := 1; i < len(nums)-1; i++ {
		solution[i] = productOfAll(nums[:i]) * productOfAll(nums[i+1:])
	}
	if len(nums) > 1 {
		solution[len(nums)-1] = productOfAll(nums[:len(nums)-1])
	}

	return solution
}

func productOfAll(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		res *= nums[i]
	}

	return res
}
