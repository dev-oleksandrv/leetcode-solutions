package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	ans, current, left := 0, 0, 0
	for right := 0; right < len(s); right++ {
		current += abs(s[right], t[right])
		if current <= maxCost {
			if right - left + 1 > ans {
				ans = right - left + 1
			}
			continue;
		}
		current -= abs(s[left], t[left])
		left++
	}
	return ans
}
func abs(s, t byte) int {
	sm := int(s) - int(t)
	if sm < 0 {
		sm *= -1
	}
	return sm
}

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
}