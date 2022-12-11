package main

import (
	u "advent-of-code-2022/utilities"
	"strconv"
)

func main() {
	scanner := u.ScannerFromFile("day1/input")
	var threeMost [3]int
	current := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for i := 0; i < len(threeMost); i++ {
				if current > threeMost[i] {
					threeMost[i] = current
					break
				}
			}
			current = 0
			continue
		}
		tmp, _ := strconv.Atoi(line)
		current += tmp
	}
	most := 0
	for i := 0; i < len(threeMost); i++ {
		most += threeMost[i]
	}
	println(most)
}
