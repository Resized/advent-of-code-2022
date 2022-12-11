package main

import (
	u "advent-of-code-2022/utilities"
	"unicode"
)

func main() {
	scanner := u.ScannerFromFile("day3/input")
	prioritiesSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		prioritiesSum += priorities(line)
	}
	println(prioritiesSum)
}

func priorities(line string) int {
	leftLine := line[0 : len(line)/2]
	rightLine := line[len(line)/2:]
	c := findCommonChar(leftLine, rightLine)
	return calcPriority(c)
}

func calcPriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c - 'A' + 27)
	} else {
		return int(c - 'a' + 1)
	}
}

func findCommonChar(line1 string, line2 string) rune {
	var charMap = make(map[byte]bool)
	for i := 0; i < len(line1); i++ {
		charMap[line1[i]] = true
	}
	for i := 0; i < len(line2); i++ {
		if charMap[line2[i]] {
			return rune(line2[i])
		}
	}
	return ' '
}
