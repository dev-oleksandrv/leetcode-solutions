package main

import "fmt"

// Time O(n), Space O(n)
// Space could be better, buy using integers instead of slices
func pivotIndex(nums []int) int {
	sumLeft := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		sumLeft[i] = sumLeft[i - 1] + nums[i - 1]
		sumRight[i] = sumRight[i - 1] + nums[len(nums) - i]
	}
	for i := 0; i < len(sumLeft); i++ {
		if sumLeft[i] == sumRight[len(sumRight) - i - 1] {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2,1,-1}))
}