package main

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}


func main() {
	fmt.Println(getCommon([]int{1,2,3}, []int{2,4}))
	fmt.Println(getCommon([]int{1,2,3,6}, []int{2,3,4,5}))
	fmt.Println(getCommon([]int{1,3,5,7}, []int{2,3,4,6,8,10,12,14,16}))
}