package main

import (
	"fmt"
	"math"
)

func findMaxLength(nums []int) int {
	occs, left, ans := map[int]int{0: 0, 1: 0}, 0, 0
	occs[nums[0]]++
	for right := 1; right < len(nums); right++ {
		occs[nums[right]]++
		for math.Abs(float64(occs[0] - occs[1])) > 1 {
			occs[nums[left]]--
			left++
		} 
		if occs[0] == occs[1] && occs[0] + occs[1] > ans {
			ans = occs[0] + occs[1]
		}
	}
	return ans
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 1, 0}))
	fmt.Println(findMaxLength([]int{0,0,0,1,1,1,0}))
}