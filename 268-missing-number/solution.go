package main

import "fmt"

func missingNumber(nums []int) int {
	expectedSum, actualSum := (len(nums) * (len(nums) + 1)) / 2, 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1}))
} 