package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	ocr := map[rune]int{}
	for _, ch := range s {
		ocr[ch]++
	}
	trg := 0
	for _, val := range ocr {
		if trg == 0 {
			trg = val
			continue
		}

		if val != trg {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
	fmt.Println(areOccurrencesEqual("aaabb"))
}