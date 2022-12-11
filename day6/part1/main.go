package main

import (
	u "advent-of-code-2022/utilities"
)

func main() {
	scanner := u.ScannerFromFile("day6/input")
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = getStartOfPacket(line, 4)
	}

	println(result)
}

func getStartOfPacket(line string, length int) int {
	uniqueCharsMap := make(map[rune]bool)
	for i := length - 1; i < len(line); i++ {
		allCharsUnique := true
		for j := i - (length - 1); j <= i; j++ {
			if _, ok := uniqueCharsMap[rune(line[j])]; ok {
				allCharsUnique = false
				uniqueCharsMap = make(map[rune]bool)
				break
			}
			uniqueCharsMap[rune(line[j])] = true
		}
		if allCharsUnique {
			return i + 1
		}
	}
	return -1
}
