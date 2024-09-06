package main

import "fmt"


func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func maxVowels(s string, k int) int {
	current := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			current++
		}
	}
	ans := current
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			current++
		}
		if isVowel(s[i - k]) {
			current--
		}
		if current > ans {
			ans = current
		}
		if ans == k {
			return k
		}
	}
	return ans
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
}