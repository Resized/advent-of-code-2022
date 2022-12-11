package main

import (
	"strconv"
	"strings"
)
import u "advent-of-code-2022/utilities"

func main() {
	data := u.ReadFile("day1/input")
	most := 0
	current := 0
	s := string(data)

	for _, line := range strings.Split(s, "\r\n") {
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
