package main

import (
	u "advent-of-code-2022/utilities"
	"strconv"
)

func main() {
	scanner := u.ScannerFromFile("day1/input")
	most := 0
	current := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if current > most {
				most = current
			}
			current = 0
			continue
		}
		tmp, _ := strconv.Atoi(line)
		current += tmp
	}

	println(most)
}
