package main

import (
	u "advent-of-code-2022/utilities"
)

type RPS int

const (
	rock     RPS = 1
	paper        = 2
	scissors     = 3
)

var vsScoreMap = map[RPS]map[RPS]int{
	rock: {
		rock:     3 + int(rock),
		paper:    6 + int(paper),
		scissors: 0 + int(scissors),
	},
	paper: {
		rock:     0 + int(rock),
		paper:    3 + int(paper),
		scissors: 6 + int(scissors),
	},
	scissors: {
		rock:     6 + int(rock),
		paper:    0 + int(paper),
		scissors: 3 + int(scissors),
	},
}

var charMap = map[byte]RPS{
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

func vs(a, b RPS) int {
	return vsScoreMap[a][b]
}

func main() {
	scanner := u.ScannerFromFile("day2/input")
	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		score += vs(charMap[line[0]], charMap[line[2]])
	}
	println(score)
}
