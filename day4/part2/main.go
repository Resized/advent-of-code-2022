package main

import (
	u "advent-of-code-2022/utilities"
	"strconv"
	"strings"
)

func main() {
	scanner := u.ScannerFromFile("day4/input")
	numOverlapping := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		leftPair := line[0]
		rightPair := line[1]
		numOverlapping += overlapping(leftPair, rightPair)
	}
	println(numOverlapping)
}

func overlapping(pair string, pair2 string) int {
	left := strings.Split(pair, "-")
	right := strings.Split(pair2, "-")
	leftStart, _ := strconv.Atoi(left[0])
	leftEnd, _ := strconv.Atoi(left[1])
	rightStart, _ := strconv.Atoi(right[0])
	rightEnd, _ := strconv.Atoi(right[1])
	if (leftStart >= rightStart && leftStart <= rightEnd) ||
		(leftStart <= rightStart && leftEnd >= rightStart) {
		return 1
	}
	return 0
}
