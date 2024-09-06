package main

import "fmt"

func reverseOnlyLetters(s string) string {
    result := []rune(s)
	left, right := 0, len(result) - 1
	for left < right {
		if !isLetter(result[right]) {
			right--
			continue
		}
		if !isLetter(result[left]) {
			left++
			continue
		}
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return string(result)
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
	fmt.Println(reverseOnlyLetters("ab-cd"))
}