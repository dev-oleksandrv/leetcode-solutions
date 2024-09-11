package main

import (
	"fmt"
	"math"
)

func maxNumberOfBalloons(text string) int {
	occs := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for _, ch := range text {
		if _, ok := occs[rune(ch)]; ok {
			occs[rune(ch)]++
		}
	}
	ans := float64(occs['b'])
	for k, v := range occs {
		fv := float64(v)
		if k == 'l' || k == 'o' {
			fv /= 2
		}
		ans = math.Min(fv, ans)
	}
	return int(ans)
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("leetcode"))
	fmt.Println(maxNumberOfBalloons("balon"))
}