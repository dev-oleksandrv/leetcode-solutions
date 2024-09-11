package main

import "fmt"

type NumArray struct {
	prefix []int
}


func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums) + 1)

	for i := 0; i < len(nums); i++ {
		prefix[i + 1] = prefix[i] + nums[i]
	}

	return NumArray{ prefix }
}


func (numArray *NumArray) SumRange(left int, right int) (sum int) {
	return numArray.prefix[right + 1] - numArray.prefix[left]
}

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}