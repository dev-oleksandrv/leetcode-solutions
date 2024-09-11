package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	winners := map[int]int{}
	losers := map[int]int{}
	for _, m := range matches {
		winners[m[0]]++
		losers[m[1]]++
	}
	w, l := []int{}, []int{}
	for k := range winners {
		if losers[k] == 0 {
			w = append(w, k)
		}
	}
	for k, v := range losers {
		if v == 1 {
			l = append(l, k)
		}
	}
	sort.Ints(w)
	sort.Ints(l)
	return [][]int{w,l}
}

func main() {
	pairs := [][]int{
        {1, 3},
        {2, 3},
        {3, 6},
        {5, 6},
        {5, 7},
        {4, 5},
        {4, 8},
        {4, 9},
        {10, 4},
        {10, 9},
    }

    fmt.Println(findWinners(pairs)) // [[1,2,10], [4,5,7,8]]
}