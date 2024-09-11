package main

import (
	"fmt"
	"sort"
)

func intersection(nums [][]int) []int {
	ocr := map[int]int{}
	for _, row := range nums {
		for _, num := range row {
			ocr[num]++
		}
	}
	res := []int{}
	for key := range ocr {
		if ocr[key] == len(nums) {
			res = append(res, key)
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	fmt.Println(intersection([]([]int){[]int{3,1,2,4,5}, []int{1,2,3,4}, []int{3,4,5,6}}))
}