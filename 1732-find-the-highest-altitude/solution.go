package main

import "fmt"

func largestAltitude(gain []int) int {
	ans, current := 0, 0
	for i := 0; i < len(gain); i++ {
		current += gain[i]
		if current > ans {
			ans = current
		}
	}
	return ans
}

func main() {
	fmt.Println(largestAltitude([]int{-5,1,5,0,-7}))
	fmt.Println(largestAltitude([]int{-4,-3,-2,-1,4,3,2}))
}