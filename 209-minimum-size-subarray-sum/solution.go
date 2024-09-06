package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	ans, left, current := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		current += nums[right]
		for current >= target {
			tmp := right - left + 1
			if ans == 0 || tmp < ans {
				ans = tmp
			}
			current -= nums[left]
			left++
		}
	}
	return ans
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1,4,4}))
	fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1}))
}