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
	line1 = u.SortString(line1)
	line2 = u.SortString(line2)
	line1Index := 0
	line2Index := 0

	for {
		if line1[line1Index] == line2[line2Index] {
			return rune(line1[line1Index])
		} else if line1[line1Index] < line2[line2Index] {
			line1Index++
		} else {
			line2Index++
		}
	}
}
