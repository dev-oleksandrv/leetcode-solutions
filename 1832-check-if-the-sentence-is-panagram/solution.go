package main

import "fmt"

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	alphabet := map[byte]bool{}
	for i := 0; i < len(sentence); i++ {
		if !alphabet[sentence[i]] {
			alphabet[sentence[i]] = true
		}
		if len(alphabet) == 26 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}