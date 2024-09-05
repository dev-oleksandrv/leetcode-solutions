package main

import "fmt"

func moveZeroes(nums []int) []int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
	return nums
}


func main() {
	fmt.Println(moveZeroes([]int{0,1,0,3,12}))
	fmt.Println(moveZeroes([]int{0}))
	fmt.Println(moveZeroes([]int{0,0,0,4,0,3,0,12,0,444,0}))
	// 0 1 0 3 12, r=0, l=0
	// 1 0 0 3 12, r=3, l=1
	// 1 3 0 0 12, r=4, l=2
	// 1 3 12 0 0

}