package main

import "fmt"

func countElements(arr []int) (out int) {
	numsMap := map[int]bool{}
	for _, num := range arr {
		numsMap[num] = true;
	}
	for _, num := range arr {
		if numsMap[num + 1] {
			out++
		}
	}
	return 
}

func main() {
	fmt.Println(countElements([]int{1, 2, 3}))
	fmt.Println(countElements([]int{1,1,3,3,5,5,7,7}))
	fmt.Println(countElements([]int{1,3,2,3,5,0}))
	fmt.Println(countElements([]int{1,1,2}))
}