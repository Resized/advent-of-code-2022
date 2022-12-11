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

var winMap = map[RPS]RPS{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

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
}

func vs(a, b RPS) int {
	return vsScoreMap[a][b]
}

func main() {
	scanner := u.ScannerFromFile("day2/input")
	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		c1 := charMap[line[0]]
		c2 := loseDrawWinToRPS(c1, line[2])
		score += vs(c1, c2)
	}
	println(score)
}

func loseDrawWinToRPS(rps RPS, loseDrawWin uint8) RPS {
	switch loseDrawWin {
	case 'X': // lose
		return winMap[rps]
	case 'Y': // draw
		return rps
	case 'Z': // win
		return winMap[winMap[rps]]
	}
	panic("Invalid input")
}
