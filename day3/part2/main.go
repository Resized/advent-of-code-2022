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
	line1 = u.SortString(line1)
	line2 = u.SortString(line2)
	line3 = u.SortString(line3)
	line1Index := 0
	line2Index := 0
	line3Index := 0

	for {
		if line1[line1Index] == line2[line2Index] && line2[line2Index] == line3[line3Index] {
			return rune(line1[line1Index])
		} else if line1[line1Index] < line2[line2Index] {
			line1Index++
		} else if line2[line2Index] < line3[line3Index] {
			line2Index++
		} else {
			line3Index++
		}
	}
}
