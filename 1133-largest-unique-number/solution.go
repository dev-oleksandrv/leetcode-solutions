package main

import "fmt"

func largestUniqueNumber(nums []int) int {
	occs := map[int]int{}
	for _, num := range nums {
		occs[num]++
	}
	ans := -1
	for num, occ := range occs {
		if occ == 1 && num > ans {
			ans = num
		}
	}
	return ans
}

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
}