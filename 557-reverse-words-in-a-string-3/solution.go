package main

import "fmt"

func reverseWords(s string) string {
	crs, left := []byte(s), 0
	for right := 0; right < len(crs); right++ {
		if crs[right] == ' ' {
			left = right + 1
			continue;
		}
		if right == len(crs) - 1 || crs[right + 1] == ' ' {
			for i := 0; i < (right - left + 1) / 2; i++ {
				crs[left + i], crs[right - i] = crs[right - i], crs[left + i]
			}
		}
	} 
	return string(crs)
}

func main() {
	// fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("LeetCode"))
	fmt.Println(reverseWords("Mr Ding"))
}