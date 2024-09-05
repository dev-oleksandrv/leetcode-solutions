package main

import "fmt"

func reversePrefix(word string, ch byte) string {
    rs := []byte(word)
	for i := 0; i < len(rs); i++ {
		if rs[i] != ch {
			continue
		}
		for j := 0; j <= int(i / 2); j++ {
			rs[j], rs[i - j] = rs[i - j], rs[j]
		}
		return string(rs)
	}
	return word
}


func main() {
	fmt.Println(reversePrefix("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}