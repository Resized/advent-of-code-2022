package main

import (
	u "advent-of-code-2022/utilities"
	"unicode"
)

func main() {
	scanner := u.ScannerFromFile("day3/input")
	prioritiesSum := 0

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		prioritiesSum += priorities(line1, line2, line3)
	}
	println(prioritiesSum)
}

func priorities(line1 string, line2 string, line3 string) int {
	c := findCommonChar(line1, line2, line3)
	return calcPriority(c)
}

func calcPriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c - 'A' + 27)
	} else {
		return int(c - 'a' + 1)
	}
}

func findCommonChar(line1 string, line2 string, line3 string) rune {
	var charMap = make(map[byte]bool)
	var charMap2 = make(map[byte]bool)
	for i := 0; i < len(line1); i++ {
		charMap[line1[i]] = true
	}
	for i := 0; i < len(line2); i++ {
		if charMap[line2[i]] {
			charMap2[line2[i]] = true
		}
	}
	for i := 0; i < len(line3); i++ {
		if charMap2[line3[i]] {
			return rune(line3[i])
		}
	}
	return ' '
}
