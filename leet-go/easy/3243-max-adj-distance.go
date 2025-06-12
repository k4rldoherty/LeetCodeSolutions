package main

import (
	"fmt"
)

func absDiffInt(x, y int) int {
	var d int
	if x < y {
		d = y - x
	} else {
		d = x - y
	}
	return d
}

func maxAdjacentDistance(nums []int) int {
	max := -1
	edgeDiff := absDiffInt(nums[0], nums[len(nums)-1])
	if edgeDiff > max {
		max = edgeDiff
	}
	for i := range len(nums) - 1 {
		d := absDiffInt(nums[i], nums[i+1])
		if d > max {
			max = d
		}
	}
	return max
}

func main() {
	testCases := [][]int{
		{1, 2, 4},
		{-5, -10, -5},
	}
	answers := []int{
		3, 5,
	}

	for i := range testCases {
		r := maxAdjacentDistance(testCases[i])
		if r != answers[i] {
			fmt.Printf("Failed on test %v\nExpected %v\nGot %v\n", i, answers[i], r)
			return
		}
	}

	fmt.Println("All tests passed!")
}
